package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/HengMrZ/chat_azure/internal/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adminToken, ok := body["admin_token"].(string)
	if ok {
		user, err := models.QueryUserByToken(models.GlobalDB, adminToken)
		if err != nil {
			http.Error(w, "admin token not found", http.StatusUnauthorized)
			return
		}
		if user.Status != 2 {
			http.Error(w, "admin token not found", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "admin token not found", http.StatusUnauthorized)
		return
	}

	userName, okUserName := body["username"].(string)
	token, okToken := body["token"].(string)
	if okUserName && okToken {
		if userName != "" && token != "" {
			err = models.AddUser(models.GlobalDB, userName, token, 1)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
