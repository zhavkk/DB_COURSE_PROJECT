package auth

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ChechPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "ivalid request payload")
		return
	}
	//TODO
	user.PasswordHash, err = HashPassword(user.PasswordHash)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}
	err = db.CreateUser(&user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "failed to register user")
		return
	}
	token, err := GenerateJWT(user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := db.GetUserByLogin(credentials.Login)
	if err != nil || user.PasswordHash != credentials.Password { // Проверка пароля
		utils.ResponseWithError(w, http.StatusUnauthorized, "invalid login or password")
		return
	}

	// Генерация JWT
	token, err := GenerateJWT(*user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
