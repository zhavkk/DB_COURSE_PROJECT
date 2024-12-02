package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	var ServiceRequest models.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&ServiceRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := db.CreateServiceRequest(&ServiceRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ServiceRequest)
}

func GetServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ServiceRequest, err := db.GetServiceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ServiceRequest)
}

func GetServiceRequestsHandler(w http.ResponseWriter, r *http.Request) {
	serviceRequests, err := db.GetAllServiceRequests()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceRequests)
}

func UpdateServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var ServiceRequest models.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&ServiceRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ServiceRequest.ID = id
	err = db.UpdateServiceRequest(&ServiceRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ServiceRequest)
}

func DeleteServiceRequestHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DeleteServiceRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
