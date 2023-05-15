package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/HengMrZ/chat_azure/internal/config"
	"github.com/HengMrZ/chat_azure/internal/models"
	"github.com/HengMrZ/chat_azure/internal/pkg"
	"github.com/sirupsen/logrus"
)

func HandleCompletions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		handleOptions(w, r)
		return
	}

	var path string
	switch r.URL.Path {
	case "/v1/chat/completions":
		path = "chat/completions"
	case "/v1/completions":
		path = "completions"
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	var body map[string]interface{}

	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil && err != io.EOF {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	token, willReturn := checkUserToken(r, w)
	if willReturn {
		return
	}

	// logrus.Infof("req token:%v", token)
	modelName := body["model"].(string)
	deployName, ok := config.GlobalCfg.Mapper[modelName]
	if !ok {
		deployName = "firstGPT"
	}

	if deployName == "" {
		http.Error(w, "Missing model mapper", http.StatusForbidden)
		return
	}
	fetchAPI := fmt.Sprintf("https://%s.openai.azure.com/openai/deployments/%s/%s?api-version=%s",
		config.GlobalCfg.ResourceName, deployName, path, config.GlobalCfg.ApiVersion)

	authKey := r.Header.Get("Authorization")
	if authKey == "" {
		http.Error(w, "Not allowed", http.StatusForbidden)
		return
	}
	bodyBts, _ := json.Marshal(body)
	resp, err := pkg.Post(fetchAPI, bodyBts, map[string]string{
		"Content-Type": "application/json",
		"api-key":      strings.TrimPrefix(authKey, "Bearer "),
	})
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	reqTokens := calcuReqTokens(body)
	// logrus.Infof("reqToken:%v", reqTokens)
	if body["stream"] == false {
		for k, v := range resp.Header {
			w.Header().Set(k, strings.Join(v, ", "))
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		return
	}

	var buf bytes.Buffer
	rdr := io.TeeReader(resp.Body, &buf)
	stream(rdr, w)

	rspTokens := calcuRspTokens(&buf)
	// logrus.Infof("rspTokens:%v", rspTokens)
	for i := 0; i < 3; i++ {
		err = models.AddCount(models.GlobalDB, token, reqTokens+rspTokens)
		if err != nil {
			logrus.Error(err)
			continue
		} else {
			return
		}
	}
}

func checkUserToken(r *http.Request, w http.ResponseWriter) (string, bool) {
	queryParams := r.URL.Query()
	token := queryParams.Get("token")
	if token == "" {
		http.Error(w, "token not found in URL", http.StatusForbidden)
		return "", true
	}
	user, err := models.QueryUserByToken(models.GlobalDB, token)
	if err != nil {
		http.Error(w, "user does not exist", http.StatusForbidden)
		return "", true
	}
	if user.Status == 0 {
		http.Error(w, "user is not valid", http.StatusForbidden)
		return "", true
	}
	return token, false
}

func stream(readable io.Reader, w http.ResponseWriter) {
	buf := make([]byte, 1024)
	for {
		n, err := readable.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if n > 0 {
			w.Write(buf[:n])
		}
	}
}
