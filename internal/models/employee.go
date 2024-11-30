package models

type Employee struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Qualification string `json:"qualification"`
	Schedule      string `json:"schedule"`
	Contact_info  string `json:"contact_info"`
	UserID        int64  `json:"user_id"`
}
