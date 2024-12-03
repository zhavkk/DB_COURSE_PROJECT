package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"dbproject/internal/utils"
	"encoding/json"
	"net/http"
)

// GetUsersHandler возвращает всех пользователей
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to retrieve users: "+err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, users)
}

// GetUserHandler возвращает пользователя по ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := db.GetUserByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, user)
}

// CreateUserHandler создает нового пользователя
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация user
	if err := validate.Struct(user); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	if err := db.CreateUser(&user); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusCreated, user)
}

// UpdateUserHandler обновляет существующего пользователя
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Валидация user
	if err := validate.Struct(user); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	user.ID = id

	if err := db.UpdateUser(&user); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to update user: "+err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, user)
}

// DeleteUserHandler удаляет пользователя по ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r, "id")
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteUser(id); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to delete user: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
