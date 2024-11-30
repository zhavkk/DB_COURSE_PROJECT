package db

import (
	"database/sql"
	"dbproject/internal/models"
)

func CreateUser(user *models.User) error {
	query := `INSERT INTO users (role_id,login,password_hash) VALUES ($1,$2,$3) RETURNING id`
	err := DB.QueryRow(query, user.RoleID, user.Login, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return err
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

func UpdateUser(user *models.User) error {
	query := `UPDATE users SET role_id =$1 , login =$2,password_hash =$3 WHERE id = $4`
	_, err := DB.Exec(query, user.RoleID, user.Login, user.PasswordHash)
	return err
}

func DelereUser(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}
