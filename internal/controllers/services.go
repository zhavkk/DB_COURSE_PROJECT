package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"
)

// GetServicesHandler возвращает список всех услуг
func GetServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, err := db.GetAllServices()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to retrieve services: "+err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, services)
}

// GetServiceHandler возвращает услугу по ID
func GetServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	service, err := db.GetServiceByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving service: "+err.Error())
		return
	}

	if service == nil {
		utils.ResponseWithError(w, http.StatusNotFound, "Service not found")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, service)
}

// CreateServiceHandler создает новую услугу
func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация service
	if err := validate.Struct(service); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	if err := db.CreateService(&service); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to create service: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusCreated, service)
}

// UpdateServiceHandler обновляет существующую услугу
func UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация service
	if err := validate.Struct(service); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	service.ID = id

	if err := db.UpdateService(&service); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to update service: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, service)
}

// DeleteServiceHandler удаляет услугу по ID
func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteService(id); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to delete service: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
