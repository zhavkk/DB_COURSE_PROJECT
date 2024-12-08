// internal/auth/auth.go
package auth

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUserInput структура для входных данных регистрации
type RegisterUserInput struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   int    `json:"role_id" validate:"required"`
}

// LoginCredentials структура для входных данных логина
type LoginCredentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterUser обработчик для регистрации пользователя
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received /register request")
	var input RegisterUserInput

	// Декодирование JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Валидация входных данных
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		log.Printf("Validation failed: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Проверка, существует ли пользователь с таким логином
	existingUser, err := db.GetUserByLogin(input.Login)
	if err != nil {
		log.Printf("Database error: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if existingUser != nil {
		log.Printf("User already exists: %s", input.Login)
		utils.ResponseWithError(w, http.StatusBadRequest, "User with this login already exists")
		return
	}

	// Хеширование пароля
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	// Создание нового пользователя
	newUser := &models.User{
		Login:        input.Login,
		PasswordHash: hashedPassword,
		RoleID:       int64(input.RoleID),
	}

	err = db.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	// Генерация JWT токена
	token, err := GenerateJWT(*newUser)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Отправка успешного ответа с токеном
	log.Printf("User registered successfully: %s", input.Login)
	utils.ResponseWithJson(w, http.StatusOK, map[string]string{"token": token})
}

// LoginUser обработчик для логина пользователя
func LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received /login request")
	var credentials LoginCredentials

	// Декодирование JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Валидация входных данных
	validate := validator.New()
	err = validate.Struct(credentials)
	if err != nil {
		log.Printf("Validation failed: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Получение пользователя из базы данных
	user, err := db.GetUserByLogin(credentials.Login)
	if err != nil {
		log.Printf("Database error: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if user == nil {
		log.Printf("User not found: %s", credentials.Login)
		utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid login or password")
		return
	}

	// Проверка пароля
	if !CheckPasswordHash(credentials.Password, user.PasswordHash) {
		log.Printf("Invalid password for user: %s", credentials.Login)
		utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid login or password")
		return
	}

	// Генерация JWT токена
	token, err := GenerateJWT(*user)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	log.Printf("User logged in successfully: %s", credentials.Login)
	utils.ResponseWithJson(w, http.StatusOK, map[string]interface{}{
		"token":   token,
		"user_id": user.ID, // Добавляем user_id в ответ
	})
}

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash проверяет соответствие пароля его хешу
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
