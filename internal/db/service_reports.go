package db

import (
	"database/sql"
	"dbproject/internal/models"
	"log"
)

// CreateServiceReport - создание нового отчета о выполненной услуге
func CreateServiceReport(report *models.ServiceReport) error {
	query := `INSERT INTO service_reports (request_id, report_text, feedback) 
              VALUES ($1, $2, $3) RETURNING id`
	err := DB.QueryRow(query, report.RequestID, report.ReportText, report.Feedback).Scan(&report.ID)
	if err != nil {
		return err
	}
	log.Printf("succesful created service report for %d", report.ID)
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

// GetAllServiceReports - получает все отчеты о выполнении услуг
func GetAllServiceReports() ([]models.ServiceReport, error) {
	// SQL-запрос для получения всех отчетов
	query := `SELECT sr.id, sr.request_id, sr.report_text, sr.feedback, t1.service_type
	FROM service_reports sr
	JOIN (
    	SELECT sr2.id, s.service_type
    	FROM service_requests sr2
    	JOIN services s ON sr2.service_id = s.id  -- Указываем условие объединения
	) AS t1 ON sr.request_id = t1.id;
	`

	// Выполняем запрос
	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Считываем все строки и создаем список отчетов
	var reports []models.ServiceReport
	for rows.Next() {
		var report models.ServiceReport
		err := rows.Scan(&report.ID, &report.RequestID, &report.ReportText, &report.Feedback, &report.ServiceType)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		reports = append(reports, report)
	}

	// Проверка на ошибки после чтения всех строк
	if err := rows.Err(); err != nil {
		log.Println("Error during row iteration:", err)
		return nil, err
	}

	return reports, nil
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
