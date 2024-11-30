package db

import (
	"database/sql"
	"dbproject/internal/models"
)

func CreateEmployee(employee *models.Employee) error {
	query := `INSERT INTO employees (name,qualification,schedule,contact_info,user_id) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := DB.QueryRow(query, employee.Name, employee.Qualification, employee.Schedule, employee.Contact_info, employee.UserID).Scan(&employee.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetEmployeeByID(id int64) (*models.Employee, error) {
	query := `SELECT * FROM employees WHERE id = $1`
	row := DB.QueryRow(query, id)

	var employee models.Employee
	err := row.Scan(&employee.ID, &employee.Name, &employee.Qualification, &employee.Schedule, &employee.Contact_info, &employee.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //no empl in db
		}

		return nil, err
	}

	return &employee, nil
}

func UpdateEmployee(employee *models.Employee) error {
	query := `UPDATE employees SET name = $1, qualification = $2,schedule = $3,contact_info = $4,user_id = $5 WHERE id = $6`
	_, err := DB.Exec(query, employee.Name, employee.Qualification, employee.Schedule, employee.Contact_info, employee.UserID)
	return err
}

func DelereEmployee(id int64) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}
