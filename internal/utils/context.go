package utils

import (
	"context"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const UserRoleKey contextKey = "user_role"

// Функция для получения данных пользователя из контекста
func GetUserFromContext(ctx context.Context) (int64, int64) {
	userID, _ := ctx.Value(UserIDKey).(int64)
	role, _ := ctx.Value(UserRoleKey).(int64)
	return userID, role
}

// Функция для установки данных пользователя в контекст
func SetUserContext(ctx context.Context, userID, role int64) context.Context {
	// Нужно дважды вызвать WithValue и правильно присваивать результат
	ctx = context.WithValue(ctx, UserIDKey, userID)
	ctx = context.WithValue(ctx, UserRoleKey, role)
	return ctx
}
