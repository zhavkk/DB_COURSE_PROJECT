package auth

import (
	"dbproject/internal/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("mishanyaBOSS")

type Claims struct { // struct that will be used for perfoming data to JWT token
	UserID int64 `json:"user_id"`
	Role   int64 `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) { // generate JWT token for our user
	claims := &Claims{
		UserID: user.ID,
		Role:   user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, errors.New("Invalid or expired token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil

}
