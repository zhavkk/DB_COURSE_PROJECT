package common

import (
	"github.com/golang-jwt/jwt"
)

// Claims представляет структуру JWT claims
type Claims struct {
	UserID int64 `json:"user_id"`
	Role   int64 `json:"role"`
	jwt.StandardClaims
}
