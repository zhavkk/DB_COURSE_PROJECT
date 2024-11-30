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

// DeleteServiceRequestEmployee - удаление связи между заявкой и сотрудником
func DeleteServiceRequestEmployee(requestID, employeeID int64) error {
	query := `DELETE FROM service_request_employees WHERE request_id = $1 AND employee_id = $2`
	_, err := DB.Exec(query, requestID, employeeID)
	return err
}
