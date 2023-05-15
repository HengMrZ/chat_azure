package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/HengMrZ/chat_azure/internal/models"
)

func QueryUser(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, ok := body["token"].(string)
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := models.QueryUserByToken(models.GlobalDB, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rsp := make(map[string]any)
	rsp["username"] = user.Username
	rsp["count"] = user.Count
	rsp["status"] = user.Status
	rspBts, _ := json.Marshal(rsp)

	w.WriteHeader(http.StatusOK)
	w.Write(rspBts)
}
