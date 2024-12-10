package db

import "dbproject/internal/models"

// CreateServiceRequestEmployee - создание связи между заявкой и сотрудником
func CreateServiceRequestEmployee(requestID, employeeID int64) error {
	query := `INSERT INTO service_request_employees (request_id, employee_id) VALUES ($1, $2)`
	_, err := DB.Exec(query, requestID, employeeID)
	return err
}

// GetServiceRequestEmployees - получение списка сотрудников для заявки
func GetServiceRequestEmployees(requestID int64) ([]models.ServiceRequestEmployee, error) {
	query := `SELECT request_id, employee_id FROM service_request_employees WHERE request_id = $1`
	rows, err := DB.Query(query, requestID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.ServiceRequestEmployee
	for rows.Next() {
		var employee models.ServiceRequestEmployee
		err := rows.Scan(&employee.RequestID, &employee.EmployeeID)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func GetServiceRequestsForEmployeeId(employeeID int64) ([]models.ServiceRequestEmployee, error) {
	query := `
        SELECT sre.request_id, sre.employee_id, sr.service_type, sr.client_id, sr.status
        FROM service_request_employees sre
        JOIN (SELECT sr.id, sr.client_id, sr.service_id, sr.status, sr.request_date, s.service_type 
              FROM service_requests as sr 
              JOIN services s ON sr.service_id = s.id) sr 
        ON sre.request_id = sr.id
        WHERE sre.employee_id = $1
    `
	rows, err := DB.Query(query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.ServiceRequestEmployee
	for rows.Next() {
		var request models.ServiceRequestEmployee
		var serviceType string
		var clientID int64
		var status int // Добавляем status

		err := rows.Scan(&request.RequestID, &request.EmployeeID, &serviceType, &clientID, &status)
		if err != nil {
			return nil, err
		}

		request.ServiceType = serviceType
		request.ClientID = clientID
		request.Status = status // Устанавливаем status в структуру

		requests = append(requests, request)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}

// DeleteServiceRequestEmployee - удаление связи между заявкой и сотрудником
func DeleteServiceRequestEmployee(requestID, employeeID int64) error {
	query := `DELETE FROM service_request_employees WHERE request_id = $1 AND employee_id = $2`
	_, err := DB.Exec(query, requestID, employeeID)
	return err
}
