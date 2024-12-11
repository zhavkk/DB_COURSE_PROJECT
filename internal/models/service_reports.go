package models

type ServiceReport struct {
	ID          int64  `json:"id"`
	RequestID   int64  `json:"request_id"`
	ReportText  string `json:"report_text"` // Описание выполненной услуги
	Feedback    string `json:"feedback"`
	ServiceType string `json:"service_type"`
}
