package auth

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
} // CheckPasswordHash проверяет совпадение пароля и хеша
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RegisterUser обрабатывает регистрацию нового пользователя
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	// Валидация входных данных
	if err := validate.Struct(input); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	// Проверка существования пользователя с таким логином
	existingUser, err := db.GetUserByLogin(input.Login)
	if err != nil && err.Error() != "record not found" {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if existingUser != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "User with this login already exists")
		return
	}

	// Хеширование пароля
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to hash password: "+err.Error())
		return
	}

	// Создание пользователя
	user := models.User{
		Login:        input.Login,
		PasswordHash: hashedPassword,
		RoleID:       input.RoleID,
	}

	if err := db.CreateUser(&user); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to register user: "+err.Error())
		return
	}

	// Генерация JWT
	token, err := GenerateJWT(user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

// LoginUser обрабатывает вход пользователя
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	// Валидация входных данных
	if err := validate.Struct(credentials); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	// Получение пользователя по логину
	user, err := db.GetUserByLogin(credentials.Login)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if user == nil {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid login or password")
		return
	}

	// Проверка пароля
	if !CheckPasswordHash(credentials.Password, user.PasswordHash) {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid login or password")
		return
	}

	// Генерация JWT
	token, err := GenerateJWT(*user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
