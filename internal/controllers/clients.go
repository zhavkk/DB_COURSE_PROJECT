package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := db.CreateClient(&client)

	if err != nil {
		http.Error(w, "Failed  ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)

}

func GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := db.GetAllClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func GetClientHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	client, err := db.GetClientByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if client == nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// UpdateClient - обработчик для обновления информации о клиенте
func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var client models.Client
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client.ID = id

	err = db.UpdateClient(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// DeleteClient - обработчик для удаления клиента
func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = db.DeleteClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
