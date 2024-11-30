package models

type User struct {
	ID           int64  `json:"id"`
	RoleID       int64  `json:"role_id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}
