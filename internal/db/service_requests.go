package db

import (
	"database/sql"
	"dbproject/internal/models"
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

// DeleteServiceRequest - удаление заявки по ID
func DeleteServiceRequest(id int64) error {
	query := `DELETE FROM service_requests WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}

// GetAllServiceRequests - получение всех заявок
func GetAllServiceRequests() ([]models.ServiceRequest, error) {
	query := `SELECT id, client_id, service_id, status, request_date, completion_date FROM service_requests`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.ServiceRequest
	for rows.Next() {
		var request models.ServiceRequest
		err := rows.Scan(&request.ID, &request.ClientID, &request.ServiceID, &request.Status, &request.RequestDate, &request.CompletionDate)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}
