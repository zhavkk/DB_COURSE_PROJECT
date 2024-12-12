package db

import (
	"database/sql"
	"dbproject/internal/models"
)

func CreateClient(client *models.Client) error {
	query := `INSERT INTO clients (name, birth_date, address, medical_needs, preferences, user_id) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := DB.QueryRow(query, client.Name, client.BirthDate, client.Address, client.MedicalNeeds, client.Preferences, client.UserID).Scan(&client.ID)
	if err != nil {
		return err

	}

	return nil
}

// Предположим, что у вас есть соединение с базой данных db

func GetClientIdFromUserId(userId int64) (int64, error) {
	var clientId int64
	query := "SELECT id FROM clients WHERE user_id = $1"
	err := DB.QueryRow(query, userId).Scan(&clientId)
	if err != nil {
		return 0, err // Возвращаем ошибку, если клиент не найден
	}
	return clientId, nil
}

func GetClientByID(id int64) (*models.Client, error) {
	query := `SELECT id,name,birth_date,address,medical_needs,preferences,user_id FROM clients WHERE id = $1`
	row := DB.QueryRow(query, id)

	var client models.Client
	err := row.Scan(&client.ID, &client.Name, &client.BirthDate, &client.Address, &client.MedicalNeeds, &client.Preferences, &client.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //no client in bd
		}

		return nil, err
	}
	return &client, nil
}

func GetAllClients() ([]models.Client, error) {
	query := `SELECT * FROM clients`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clients []models.Client

	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID, &client.Name, &client.BirthDate, &client.Address, &client.MedicalNeeds, &client.Preferences, &client.UserID)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	// Проверяем, были ли ошибки при переборе строк
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil

}

// Обновление данных клиента
func UpdateClient(clientId string, updatedClient models.Client) error {
	query := `UPDATE clients SET name = $1, birth_date = $2, address = $3, medical_needs = $4 WHERE id = $5`
	_, err := DB.Exec(query, updatedClient.Name, updatedClient.BirthDate, updatedClient.Address, updatedClient.MedicalNeeds, clientId)
	return err
}

func DeleteClient(id int64) error {
	query := `DELETE FROK clients WHERE id = $1`
	_, err := DB.Exec(query, id)
	return err
}
