package api

import (
	"encoding/json"
	"net/http"

	"github.com/HengMrZ/chat_azure/internal/config"
)

func HandleModels(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Object string                   `json:"object"`
		Data   []map[string]interface{} `json:"data"`
	}{
		Object: "list",
		Data:   []map[string]interface{}{},
	}

	for key := range config.GlobalCfg.Mapper {
		model := map[string]interface{}{
			"id":       key,
			"object":   "model",
			"created":  1677610602,
			"owned_by": "openai",
			"permission": []map[string]interface{}{
				{
					"id":                   "modelperm-M56FXnG1AsIr3SXq8BYPvXJA",
					"object":               "model_permission",
					"created":              1679602088,
					"allow_create_engine":  false,
					"allow_sampling":       true,
					"allow_logprobs":       true,
					"allow_search_indices": false,
					"allow_view":           true,
					"allow_fine_tuning":    false,
					"organization":         "*",
					"group":                nil,
					"is_blocking":          false,
				},
			},
			"root":   key,
			"parent": nil,
		}
		data.Data = append(data.Data, model)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
