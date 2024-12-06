package auth

import (
	"dbproject/internal/models"
	"os"
	"testing"
	"time"
)

// Инициализация jwtKey для тестов
func init() {
	os.Setenv("JWT_SECRET_KEY", "testsecret")
	InitJWTKey()
}

func TestGenerateAndVerifyJWT(t *testing.T) {
	user := models.User{
		ID:           123,
		RoleID:       2,
		Login:        "testuser",
		PasswordHash: "hash",
	}

	token, err := GenerateJWT(user)
	if err != nil {
		t.Fatalf("unexpected error generating token: %v", err)
	}

	claims, err := VerifyJWT(token)
	if err != nil {
		t.Fatalf("unexpected error verifying token: %v", err)
	}

	if claims.UserID != user.ID {
		t.Errorf("expected userID %d, got %d", user.ID, claims.UserID)
	}
	if claims.Role != user.RoleID {
		t.Errorf("expected role %d, got %d", user.RoleID, claims.Role)
	}

	// Проверяем срок жизни токена
	if claims.ExpiresAt < time.Now().Unix() {
		t.Error("token should not be expired")
	}
}
