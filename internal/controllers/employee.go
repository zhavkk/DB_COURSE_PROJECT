package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// GetEmployeesHandler возвращает список всех сотрудников
func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := db.GetAllEmployees()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving employees: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, employees)
}

// GetEmployeeByIDHandler возвращает сотрудника по ID
func GetEmployeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	employee, err := db.GetEmployeeByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error retrieving employee: "+err.Error())
		return
	}

	if employee == nil {
		utils.ResponseWithError(w, http.StatusNotFound, "Employee not found")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, employee)
}

// CreateEmployeeHandler создаёт нового сотрудника
func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var newEmployee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация сотрудника
	if err := validate.Struct(newEmployee); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	if err := db.CreateEmployee(&newEmployee); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error creating employee: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusCreated, newEmployee)
}

// UpdateEmployeeHandler обновляет данные сотрудника
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var updatedEmployee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация сотрудника
	if err := validate.Struct(updatedEmployee); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	updatedEmployee.ID = id

	if err := db.UpdateEmployee(&updatedEmployee); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error updating employee: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, updatedEmployee)
}

// DeleteEmployeeHandler удаляет сотрудника по ID
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteEmployee(id); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error deleting employee: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
