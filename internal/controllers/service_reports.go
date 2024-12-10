package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
)

// GetAllServiceReportsHandler возвращает список всех отчетов
func GetAllServiceReportsHandler(w http.ResponseWriter, r *http.Request) {
	reports, err := db.GetAllServiceReports()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving service reports: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, reports)
}

// GetServiceReportByIDHandler возвращает отчет по ID
func GetServiceReportByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	report, err := db.GetServiceReportByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving service report: "+err.Error())
		return
	}

	if report == nil {
		utils.ResponseWithError(w, http.StatusNotFound, "Service report not found")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, report)
}

// CreateServiceReportHandler создаёт новый отчет
func CreateServiceReportHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из тела запроса
	var report models.ServiceReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Получаем токен из заголовка
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Проверяем токен (это пример, вы можете добавить свою логику)
	token, err := jwt.Parse(tokenString, nil) // Здесь нужно проверить токен
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Сохраняем отчет в базе данных
	err = db.CreateServiceReport(&report)
	if err != nil {
		http.Error(w, "Failed to save report", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Report created successfully"})
}

// UpdateServiceReportHandler обновляет отчет
func UpdateServiceReportHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var report models.ServiceReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация отчета
	if err := validate.Struct(report); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	report.ID = id

	if err := db.UpdateServiceReport(&report); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error updating service report: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, report)
}

// DeleteServiceReportHandler удаляет отчет по ID
func DeleteServiceReportHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteServiceReport(id); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error deleting service report: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
