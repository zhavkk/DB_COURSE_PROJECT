package models

import (
	"time"
)

type ServiceRequest struct {
	ID             int64      `json:"id"`
	ClientID       int64      `json:"client_id"`
	ServiceID      int64      `json:"service_id"`
	Status         int        `json:"status"` // Статус заявки (0 - в процессе, 1 - завершено)
	RequestDate    time.Time  `json:"request_date"`
	CompletionDate *time.Time `json:"completion_date"` // Для возможного NULL значения
}
