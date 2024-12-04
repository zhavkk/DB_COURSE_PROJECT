package auth

import (
	"dbproject/internal/common"
	"dbproject/internal/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey []byte

// InitJWTKey устанавливает значение jwtKey
func InitJWTKey() {
	// Загружаем ключ из переменной окружения
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		panic("JWT_SECRET_KEY environment variable is not set")
	}
	jwtKey = []byte(key)
}

// GenerateJWT генерирует JWT токен для пользователя
func GenerateJWT(user models.User) (string, error) {
	if jwtKey == nil {
		return "", errors.New("jwtKey is not initialized")
	}

	claims := &common.Claims{
		UserID: user.ID,
		Role:   user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "dbproject",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// VerifyJWT проверяет валидность JWT токена и возвращает claims
func VerifyJWT(tokenString string) (*common.Claims, error) {
	if jwtKey == nil {
		return nil, errors.New("jwtKey is not initialized")
	}

	claims := &common.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
