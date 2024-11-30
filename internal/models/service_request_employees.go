package models

type ServiceRequestEmployee struct {
	RequestID  int64 `json:"request_id"`
	EmployeeID int64 `json:"employee_id"`
}
