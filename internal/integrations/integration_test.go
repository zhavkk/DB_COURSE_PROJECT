package integration

import (
	"bytes"
	"dbproject/internal/auth"
	"dbproject/internal/db"
	"dbproject/internal/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	// Установим переменные окружения, если не выставлены
	if os.Getenv("DB_HOST") == "" {
		os.Setenv("DB_HOST", "localhost")
	}
	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "postgres")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "postgres")
	}
	if os.Getenv("DB_NAME") == "" {
		os.Setenv("DB_NAME", "test_kp")
	}

	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db.InitDB(dataSourceName)

	// Если есть миграции, запустить их
	// err := RunMigrations(dataSourceName) // реализуйте функцию RunMigrations
	// if err != nil {
	//     fmt.Println("Failed to run migrations:", err)
	//     os.Exit(1)
	// }

	os.Setenv("JWT_SECRET_KEY", "testsecret")
	auth.InitJWTKey()

	code := m.Run()

	// Закрываем соединение с БД после всех тестов
	db.CloseDB()

	os.Exit(code)
}

func TestRegisterAndLogin(t *testing.T) {
	// Установим переменную окружения для JWT-секрета, если это требуется
	os.Setenv("JWT_SECRET_KEY", "testsecret")

	// Инициализируем роутер
	r := mux.NewRouter()
	routes.SetupRoutes(r)

	// Создаём тестовый сервер
	server := httptest.NewServer(r)
	defer server.Close()

	// 1. Тестируем регистрацию
	registerPayload := `{"login":"mishGuN","password":"testpass","role_id":3}`
	resp, err := http.Post(server.URL+"/register", "application/json", bytes.NewBufferString(registerPayload))
	if err != nil {
		t.Fatalf("Error making /register request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got: %d", resp.StatusCode)
	}
	var registerResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&registerResp); err != nil {
		t.Fatalf("error decoding register response: %v", err)
	}
	resp.Body.Close()

	token, ok := registerResp["token"]
	if !ok || token == "" {
		t.Fatalf("no token in register response")
	}

	// 2. Тестируем логин
	loginPayload := `{"login":"mishGuN","password":"testpass"}`
	resp, err = http.Post(server.URL+"/login", "application/json", bytes.NewBufferString(loginPayload))
	if err != nil {
		t.Fatalf("error making /login request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK on /login, got %d", resp.StatusCode)
	}
	var loginResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("error decoding login response: %v", err)
	}
	resp.Body.Close()

	loginToken, ok := loginResp["token"]
	if !ok || loginToken == "" {
		t.Fatalf("no token in login response")
	}

	// Дополнительно можно проверить, что токен при логине отличается от токена при регистрации,
	// но обычно это не обязательно, так как при регистрации мы уже возвращаем свежесгенерированный токен.
}
func TestClientsEndpoint(t *testing.T) {
	os.Setenv("JWT_SECRET_KEY", "testsecret")

	r := mux.NewRouter()
	routes.SetupRoutes(r)
	server := httptest.NewServer(r)
	defer server.Close()

	// Регистрируемся под сотрудником (role_id=2)
	registerPayload := `{"login":"employe","password":"emppass","role_id":2}`
	resp, err := http.Post(server.URL+"/register", "application/json", bytes.NewBufferString(registerPayload))
	if err != nil {
		t.Fatalf("error making register request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	var regData map[string]string
	json.NewDecoder(resp.Body).Decode(&regData)
	resp.Body.Close()

	token := regData["token"]
	if token == "" {
		t.Fatalf("no token returned on registration")
	}

	// Теперь обращаемся к /clients с этим токеном
	req, err := http.NewRequest("GET", server.URL+"/clients", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("error making /clients request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 on /clients, got %d", resp.StatusCode)
	}

	var clientsData []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&clientsData); err != nil {
		t.Fatalf("error decoding /clients response: %v", err)
	}
	resp.Body.Close()

	// Можно дополнительно проверить структуру данных, если знаете что должно вернуться
}
func TestProtectedEndpointWithoutToken(t *testing.T) {
	os.Setenv("JWT_SECRET_KEY", "testsecret")

	r := mux.NewRouter()
	routes.SetupRoutes(r)
	server := httptest.NewServer(r)
	defer server.Close()

	// Обращаемся к /clients без Authorization
	req, err := http.NewRequest("GET", server.URL+"/clients", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("error making request: %v", err)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 Unauthorized, got %d", resp.StatusCode)
	}
}
