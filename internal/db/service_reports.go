package db

import (
	"database/sql"
	"dbproject/internal/models"
)

// CreateServiceReport - создание нового отчета о выполненной услуге
func CreateServiceReport(report *models.ServiceReport) error {
	query := `INSERT INTO service_reports (request_id, report_text, feedback) 
              VALUES ($1, $2, $3) RETURNING id`
	err := DB.QueryRow(query, report.RequestID, report.ReportText, report.Feedback).Scan(&report.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetServiceReportByID - получение отчета по ID
func GetServiceReportByID(id int64) (*models.ServiceReport, error) {
	query := `SELECT id, request_id, report_text, feedback FROM service_reports WHERE id = $1`
	row := DB.QueryRow(query, id)

	var report models.ServiceReport
	err := row.Scan(&report.ID, &report.RequestID, &report.ReportText, &report.Feedback)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Нет такого отчета
		}
		return nil, err
	}

	return &report, nil
}

// UpdateServiceReport - обновление отчета
func UpdateServiceReport(report *models.ServiceReport) error {
	query := `UPDATE service_reports SET request_id = $1, report_text = $2, feedback = $3 WHERE id = $4`
	_, err := DB.Exec(query, report.RequestID, report.ReportText, report.Feedback, report.ID)
	return err
}

// DeleteServiceReport - удаление отчета по ID
func DeleteServiceReport(id int64) error {
	query := `DELETE FROM service_reports WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}
