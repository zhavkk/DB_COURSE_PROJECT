package models

type ServiceRequestEmployee struct {
	RequestID   int64  `json:"request_id"`
	EmployeeID  int64  `json:"employee_id"`
	ServiceType string `json:"service_type"`
	ClientID    int64  `json:"client_id"` // Добавляем client_id
	Status      int    `json:"status"`    // Добавляем статус

}
