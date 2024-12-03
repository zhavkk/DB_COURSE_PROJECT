package models

type User struct {
	ID           int64  `json:"id"`
	RoleID       int64  `json:"role_id"`
	Login        string `json:"login" validate:"required,min=3,max=32"`
	PasswordHash string `json:"-"`
}
