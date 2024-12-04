package db

import (
	"database/sql"
	"dbproject/internal/models"
	"errors"
	"fmt"
)

// CreateUser создает нового пользователя в базе данных
func CreateUser(user *models.User) error {
	if DB == nil {
		return errors.New("database is not initialized")
	}

	query := "INSERT INTO users (login, password_hash, role_id) VALUES ($1, $2, $3) RETURNING id"
	err := DB.QueryRow(query, user.Login, user.PasswordHash, user.RoleID).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func GetUserByID(id int64) (*models.User, error) {
	query := `SELECT id,role_id,login,password_hash FROM users WHERE id = $1`
	row := DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.RoleID, &user.Login, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user in db
		}

		return nil, err
	}

	return &user, nil
}

// GetUserByLogin возвращает пользователя по логину
func GetUserByLogin(login string) (*models.User, error) {
	if DB == nil {
		return nil, errors.New("database is not initialized")
	}

	query := "SELECT id, login, password_hash, role_id FROM users WHERE login = $1"
	row := DB.QueryRow(query, login)

	var user models.User
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.RoleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Пользователь не найден
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}
func GetAllUsers() ([]models.User, error) {
	query := `SELECT * FROM users`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		// Декодируем все поля из строки в структуру
		err := rows.Scan(&user.ID, &user.RoleID, &user.Login, &user.PasswordHash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Проверяем, были ли ошибки при переборе строк
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user *models.User) error {
	query := `UPDATE users SET role_id =$1 , login =$2,password_hash =$3 WHERE id = $4`
	_, err := DB.Exec(query, user.RoleID, user.Login, user.PasswordHash)
	return err
}

func DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}
