package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Helper function to parse ID from path
func parseIDFromPath(r *http.Request, key string) (int64, error) {
	vars := mux.Vars(r)
	idStr, exists := vars[key]
	if !exists || idStr == "" {
		return 0, fmt.Errorf("ID is required")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format")
	}
	return id, nil
}

// CreateClientHandler создаёт нового клиента
func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация клиента
	if err := validate.Struct(client); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	if err := db.CreateClient(&client); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to create client: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusCreated, client)
}

// GetClientsHandler возвращает список всех клиентов
func GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := db.GetAllClients()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving clients: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, clients)
}

func GetClientByUSERIDHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	// Преобразуем userId в int64
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Получаем client_id из базы данных
	clientId, err := db.GetClientIdFromUserId(userIdInt64)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	// Отправляем client_id в ответе
	response := map[string]int64{"client_id": clientId}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetClientHandler возвращает клиента по ID
func GetClientHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	client, err := db.GetClientByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving client: "+err.Error())
		return
	}
	if client == nil {
		utils.ResponseWithError(w, http.StatusNotFound, "Client not found")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, client)
}

// UpdateClientHandler обновляет информацию о клиенте
// UpdateClientHandler обновляет информацию о клиенте
func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientId := params["id"]

	// Проверка токена (псевдокод, нужно реализовать логику)
	token := r.Header.Get("Authorization")
	if token == "" {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Неверный токен")
		return
	}

	var updatedClient models.Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Ошибка при декодировании данных")
		return
	}

	// Обновляем клиента в базе данных
	err = db.UpdateClient(clientId, updatedClient)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Ошибка при обновлении данных клиента")
		return
	}
	log.Printf("updated for cliend id %s", clientId)
	utils.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Данные успешно обновлены"})
}

// DeleteClientHandler удаляет клиента по ID
func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteClient(id); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error deleting client: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
