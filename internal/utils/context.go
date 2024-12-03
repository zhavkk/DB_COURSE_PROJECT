package utils

import (
	"context"
	"errors"

	"dbproject/internal/common"
)

type contextKey string

const userContextKey = contextKey("user")

// GetUserFromContext извлекает UserID и Role из контекста
func GetUserFromContext(ctx context.Context) (int64, int64, error) {
	claims, ok := ctx.Value(userContextKey).(*common.Claims)
	if !ok || claims == nil {
		return 0, 0, errors.New("user not found in context")
	}
	return claims.UserID, claims.Role, nil
}

// SetUserContext устанавливает claims в контекст
func SetUserContext(ctx context.Context, claims *common.Claims) context.Context {
	return context.WithValue(ctx, userContextKey, claims)
}
