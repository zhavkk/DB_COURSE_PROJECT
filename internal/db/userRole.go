package db

import "dbproject/internal/models"

func GetAllRoles() ([]models.UserRole, error) {
	query := `SELECT id,role_name FROM user_roles`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.UserRole
	for rows.Next() {
		var role models.UserRole
		if err := rows.Scan(&role.ID, &role.RoleName); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil

}
