package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetAllServiceReports - обработчик для получения списка всех отчетов
func GetAllServiceReports(w http.ResponseWriter, r *http.Request) {
	reports, err := db.GetAllServiceReports()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Преобразуем список отчетов в JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reports)
}

// GetServiceReportByID - обработчик для получения отчета по ID
func GetServiceReportByID(w http.ResponseWriter, r *http.Request) {
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

	report, err := db.GetServiceReportByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if report == nil {
		http.Error(w, "Service report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}

// CreateServiceReport - обработчик для создания нового отчета
func CreateServiceReport(w http.ResponseWriter, r *http.Request) {
	var report models.ServiceReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateServiceReport(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(report)
}

// UpdateServiceReport - обработчик для обновления отчета
func UpdateServiceReport(w http.ResponseWriter, r *http.Request) {
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

	var report models.ServiceReport
	err = json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	report.ID = id

	err = db.UpdateServiceReport(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}

// DeleteServiceReport - обработчик для удаления отчета
func DeleteServiceReport(w http.ResponseWriter, r *http.Request) {
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

	err = db.DeleteServiceReport(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
