package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseWithError отправляет ответ с ошибкой в формате JSON
func ResponseWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}

// ResponseWithJson отправляет успешный ответ с данными в формате JSON
func ResponseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}
