package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/HengMrZ/chat_azure/internal/config"
	"github.com/HengMrZ/chat_azure/internal/pkg"
	"github.com/pandodao/tokenizer-go"
	"github.com/sirupsen/logrus"
)

func calculateTokens(input string) int {
	return tokenizer.MustCalToken(input)
}

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
	logrus.Infof("req body:%v", body)
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

	if body["stream"] == false {
		for k, v := range resp.Header {
			w.Header().Set(k, strings.Join(v, ", "))
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		return
	}
	tokenCaculator := io.Discard // 统计token
	rdr := io.TeeReader(resp.Body, tokenCaculator)
	stream(rdr, w)
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
