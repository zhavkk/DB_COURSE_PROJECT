package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateServiceRequestHandler создает новый запрос на услугу
func CreateServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	var serviceRequest models.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&serviceRequest); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация serviceRequest
	if err := validate.Struct(serviceRequest); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	if err := db.CreateServiceRequest(&serviceRequest); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to create service request: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusCreated, serviceRequest)
}

// GetServiceRequestHandler возвращает запрос на услугу по ID
func GetServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	serviceRequest, err := db.GetServiceByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving service request: "+err.Error())
		return
	}

	if serviceRequest == nil {
		utils.ResponseWithError(w, http.StatusNotFound, "Service request not found")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, serviceRequest)
}

// GetServiceRequestsHandler возвращает все запросы на услугу
func GetServiceRequestsHandler(w http.ResponseWriter, r *http.Request) {
	serviceRequests, err := db.GetAllServiceRequests()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to retrieve service requests: "+err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, serviceRequests)
}

// UpdateServiceRequestHandler обновляет существующий запрос на услугу
func UpdateServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var serviceRequest models.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&serviceRequest); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация serviceRequest
	if err := validate.Struct(serviceRequest); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	serviceRequest.ID = id

	if err := db.UpdateServiceRequest(&serviceRequest); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to update service request: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, serviceRequest)
}
func UpdateServiceRequestStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем token из заголовков
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	// Извлекаем request_id из URL пути
	vars := mux.Vars(r)
	requestID := vars["request_id"]
	if requestID == "" {
		http.Error(w, "Request ID is required", http.StatusBadRequest)
		return
	}

	// Преобразуем requestID в int64
	id, err := strconv.ParseInt(requestID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для обновления статуса
	err = db.FinishStatusServiceRequest(id)
	if err != nil {
		http.Error(w, "Error updating status", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Status updated to completed"})
}

// DeleteServiceRequestHandler удаляет запрос на услугу по ID
func DeleteServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteServiceRequest(id); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to delete service request: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
