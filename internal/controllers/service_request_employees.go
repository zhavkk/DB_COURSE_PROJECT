package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateServiceRequestEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем тело запроса
	var data struct {
		RequestID  int64 `json:"request_id"`
		EmployeeID int64 `json:"employee_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Логируем полученные данные
	log.Printf("Received request_id: %d, employee_id: %d", data.RequestID, data.EmployeeID)

	// Создаем связь между заявкой и сотрудником
	err = db.CreateServiceRequestEmployee(data.RequestID, data.EmployeeID)
	if err != nil {
		http.Error(w, "Error creating service request employee", http.StatusInternalServerError)
		return
	}

	// Логируем попытку обновления статуса
	log.Printf("Updating status for request_id: %d to 'In Progress' (status = 1)", data.RequestID)

	// Обновляем статус заявки на "В процессе" (status = 1)
	err = db.UpdateStatusServiceRequest(&models.ServiceRequest{
		ID:     data.RequestID,
		Status: 1, // В процессе
	})
	if err != nil {
		log.Printf("Error updating status for request_id: %d: %v", data.RequestID, err)
		http.Error(w, "Error updating service request status", http.StatusInternalServerError)
		return
	}

	// Логируем успешное обновление
	log.Printf("Successfully updated status for request_id: %d to 'In Progress'", data.RequestID)

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Service request employee created and status updated"))
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
func GetServiceRequestsForEmployeeIdHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем employee_id из строки запроса
	employeeID := r.URL.Query().Get("employee_id")
	if employeeID == "" {
		http.Error(w, "Employee ID is required", http.StatusBadRequest)
		return
	}

	// Преобразуем employeeID в int64 (предполагаем, что он всегда будет корректным)
	id, err := strconv.ParseInt(employeeID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	log.Printf("Request received for employee_id: %d", id)

	// Получаем список заявок для указанного сотрудника
	serviceRequests, err := db.GetServiceRequestsForEmployeeId(id)
	if err != nil {
		log.Printf("Error fetching service requests: %v", err)
	} else {
		log.Printf("Found %d service requests", len(serviceRequests))
	}
	// Возвращаем успешный ответ с данными
	utils.ResponseWithJson(w, http.StatusOK, serviceRequests)
}

// GetServiceRequestsForAdminsHandler - Обработчик для получения всех заявок для администраторов
func GetServiceRequestsForAdminsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что запрос был отправлен администратором (например, через заголовок Authorization)
	token := r.Header.Get("Authorization")
	if token == "" {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Отсутствует токен авторизации.")
		return
	}

	// Получаем все заявки для администраторов
	serviceRequests, err := db.GetServiceRequestsForAdmins()
	if err != nil {
		log.Println("Ошибка при получении заявок:", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Ошибка при получении заявок.")
		return
	}
	// Если заявки найдены, отправляем их в ответе
	utils.ResponseWithJson(w, http.StatusOK, serviceRequests)
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
