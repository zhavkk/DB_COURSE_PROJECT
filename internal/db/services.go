package db

import (
	"database/sql"
	"dbproject/internal/models"
)

// CreateService - создание новой услуги
func CreateService(service *models.Service) error {
	query := `INSERT INTO services (service_type, duration) VALUES ($1, $2) RETURNING id`
	err := DB.QueryRow(query, service.ServiceType, service.Duration).Scan(&service.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetServiceByID - получение услуги по ID
func GetServiceByID(id int64) (*models.Service, error) {
	query := `SELECT id, service_type, duration FROM services WHERE id = $1`
	row := DB.QueryRow(query, id)

	var service models.Service
	err := row.Scan(&service.ID, &service.ServiceType, &service.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Нет такой услуги
		}
		return nil, err
	}

	return &service, nil
}

// UpdateService - обновление услуги
func UpdateService(service *models.Service) error {
	query := `UPDATE services SET service_type = $1, duration = $2 WHERE id = $3`
	_, err := DB.Exec(query, service.ServiceType, service.Duration, service.ID)
	return err
}

// DeleteService - удаление услуги по ID
func DeleteService(id int64) error {
	query := `DELETE FROM services WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}

// GetAllServices - получение всех услуг
func GetAllServices() ([]models.Service, error) {
	query := `SELECT id, service_type, duration FROM services`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.ID, &service.ServiceType, &service.Duration)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}
