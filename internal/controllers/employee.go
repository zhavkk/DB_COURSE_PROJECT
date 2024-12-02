package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetAllEmployees - Обработчик для получения всех сотрудников
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := db.GetAllEmployees()
	if err != nil {
		http.Error(w, "Error retrieving employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

// GetEmployeeByID - Обработчик для получения сотрудника по ID
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	employee, err := db.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, "Error retrieving employee", http.StatusInternalServerError)
		return
	}

	if employee.ID == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

// CreateEmployee - Обработчик для создания нового сотрудника
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee models.Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = db.CreateEmployee(&newEmployee)
	if err != nil {
		http.Error(w, "Error creating employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmployee)
}

// UpdateEmployee - Обработчик для обновления данных сотрудника
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var updatedEmployee models.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedEmployee.ID = id

	err = db.UpdateEmployee(&updatedEmployee)
	if err != nil {
		http.Error(w, "Error updating employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEmployee)
}

// DeleteEmployee - Обработчик для удаления сотрудника
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteEmployee(id)
	if err != nil {
		http.Error(w, "Error deleting employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
