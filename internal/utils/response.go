package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseWithError отправляет JSON-ответ с ошибкой
func ResponseWithError(w http.ResponseWriter, code int, message string) {
	ResponseWithJson(w, code, map[string]string{"error": message})
}

// ResponseWithJSON отправляет JSON-ответ с данными
func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
