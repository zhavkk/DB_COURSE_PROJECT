package auth

import (
	"dbproject/internal/utils"
	"net/http"
	"strings"
)

func TokenVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверка наличия токена в заголовке Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		// Извлечение токена из заголовка
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Bearer token missing")
			return
		}

		// Проверка и валидация токена
		_, err := VerifyJWT(tokenString) // ParseJWT должен быть функцией для валидации токена
		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Передача запроса дальше
		next.ServeHTTP(w, r)
	})
}
