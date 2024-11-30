package models

type Service struct {
	ID          int64  `json:"id"`
	ServiceType string `json:"service_type"`
	Duration    int    `json:"duration"`
}
