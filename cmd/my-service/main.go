// main.go
package main

import (
	"dbproject/internal/auth"
	"dbproject/internal/db"
	"dbproject/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Инициализация JWT ключа
	auth.InitJWTKey()

	// Построение строки подключения к базе данных из переменных окружения
	dataSourceName := buildDataSourceName()

	// Инициализация базы данных
	db.InitDB(dataSourceName)

	// Получение порта из переменных окружения или использование дефолтного значения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Создание нового роутера
	router := mux.NewRouter()

	// Настройка маршрутов
	routes.SetupRoutes(router)

	// Запуск сервера
	log.Printf("Server is running on port %s", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// buildDataSourceName строит строку подключения из переменных окружения
func buildDataSourceName() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE") // Можно добавить SSL режим, если требуется

	if sslmode == "" {
		sslmode = "disable" // По умолчанию отключено
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}
