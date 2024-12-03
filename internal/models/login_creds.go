package models

type RegisterUserInput struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   int64  `json:"role_id" validate:"required,oneof=1 2 3"` // Пример: 1 - Админ, 2 - Сотрудник, 3 - Клиент
}

type LoginCredentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}
