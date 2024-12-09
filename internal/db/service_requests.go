package db

import (
	"database/sql"
	"dbproject/internal/models"
	"log"
)

// CreateServiceRequest - создание новой заявки на услугу
func CreateServiceRequest(request *models.ServiceRequest) error {
	query := `INSERT INTO service_requests (client_id, service_id, status, request_date, completion_date)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := DB.QueryRow(query, request.ClientID, request.ServiceID, request.Status, request.RequestDate, request.CompletionDate).Scan(&request.ID)

	if err != nil {
		return err
	}
	return nil
}

// GetServiceRequestByID - получение заявки по ID
func GetServiceRequestByID(id int64) (*models.ServiceRequest, error) {
	query := `SELECT id, client_id, service_id, status, request_date, completion_date 
              FROM service_requests WHERE id = $1`
	row := DB.QueryRow(query, id)

	var request models.ServiceRequest
	err := row.Scan(&request.ID, &request.ClientID, &request.ServiceID, &request.Status, &request.RequestDate, &request.CompletionDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Нет такой заявки
		}
		return nil, err
	}

	return &request, nil
}

// UpdateServiceRequest - обновление данных заявки
func UpdateServiceRequest(request *models.ServiceRequest) error {
	query := `UPDATE service_requests SET client_id = $1, service_id = $2, status = $3, 
              request_date = $4, completion_date = $5 WHERE id = $6`
	_, err := DB.Exec(query, request.ClientID, request.ServiceID, request.Status, request.RequestDate, request.CompletionDate, request.ID)
	return err
}
func UpdateStatusServiceRequest(request *models.ServiceRequest) error {
	// Логируем входные данные
	log.Printf("Updating status for request_id: %d, status: %d", request.ID, request.Status)

	// Выполняем SQL-запрос для обновления статуса заявки
	query := `UPDATE service_requests SET status=$2 WHERE id=$1`
	_, err := DB.Exec(query, request.ID, request.Status)
	if err != nil {
		// Логируем ошибку
		log.Printf("Error updating status for request_id: %d: %v", request.ID, err)
		return err
	}

	// Логируем успешное обновление
	log.Printf("Successfully updated status for request_id: %d", request.ID)
	return nil
}

// DeleteServiceRequest - удаление заявки по ID
func DeleteServiceRequest(id int64) error {
	query := `DELETE FROM service_requests WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}

// GetAllServiceRequests - получение всех заявок
func GetAllServiceRequests() ([]models.ServiceRequest, error) {
	query := `SELECT sr.id, sr.client_id, sr.service_id, sr.status, sr.request_date, s.service_type FROM service_requests as sr JOIN services s ON sr.service_id = s.id`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.ServiceRequest
	for rows.Next() {
		var request models.ServiceRequest
		err := rows.Scan(&request.ID, &request.ClientID, &request.ServiceID, &request.Status, &request.RequestDate, &request.ServiceType)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}
