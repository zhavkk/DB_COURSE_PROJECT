package controllers

import (
	"dbproject/internal/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateServiceRequestEmployeeHandler - создание связи между заявкой и сотрудником
func CreateServiceRequestEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из URL
	vars := mux.Vars(r)
	requestID := vars["request_id"]
	employeeID := vars["employee_id"]

	// Преобразуем из string в int64
	requestIDInt, err := strconv.ParseInt(requestID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}
	employeeIDInt, err := strconv.ParseInt(employeeID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Создаем связь между заявкой и сотрудником
	err = db.CreateServiceRequestEmployee(requestIDInt, employeeIDInt)
	if err != nil {
		http.Error(w, "Error creating service request employee", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Service request employee created"))
}

// GetServiceRequestEmployeesHandler - получение списка сотрудников для заявки
func GetServiceRequestEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем request_id из параметров URL
	vars := mux.Vars(r)
	requestID := vars["request_id"]

	// Преобразуем requestID в int64
	requestIDInt, err := strconv.ParseInt(requestID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}

	// Получаем список сотрудников для заявки
	employees, err := db.GetServiceRequestEmployees(requestIDInt)
	if err != nil {
		http.Error(w, "Error fetching service request employees", http.StatusInternalServerError)
		return
	}

	// Отправляем список сотрудников в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// DeleteServiceRequestEmployeeHandler - удаление связи между заявкой и сотрудником
func DeleteServiceRequestEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из URL
	vars := mux.Vars(r)
	requestID := vars["request_id"]
	employeeID := vars["employee_id"]

	// Преобразуем из string в int64
	requestIDInt, err := strconv.ParseInt(requestID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}
	employeeIDInt, err := strconv.ParseInt(employeeID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Удаляем связь между заявкой и сотрудником
	err = db.DeleteServiceRequestEmployee(requestIDInt, employeeIDInt)
	if err != nil {
		http.Error(w, "Error deleting service request employee", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusNoContent)
}
