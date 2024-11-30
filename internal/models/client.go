package models

type Client struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	BirthDate    string `json:"birth_date"`
	Address      string `json:"address"`
	MedicalNeeds string `json:"medical_needs"`
	Preferences  string `json:"preferences"`
	UserID       int64  `json:"user_id"`
}
