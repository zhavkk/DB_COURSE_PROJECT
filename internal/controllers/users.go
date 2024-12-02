package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input,", http.StatusBadRequest)
		return
	}

	err := db.CreateUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid user id ", http.StatusBadRequest)
		return
	}
	user, err := db.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to receive users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	user.ID = id
	err = db.UpdateUser(&user)
	if err != nil {
		http.Error(w, "failed to update user ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	err = db.DeleteUser(id)
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
