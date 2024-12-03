package auth

import (
	"dbproject/internal/utils"
	"net/http"
	"strings"
)

// TokenVerifyMiddleware проверяет наличие и валидность JWT токена
func TokenVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Bearer token missing")
			return
		}

		claims, err := VerifyJWT(tokenString)
		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		// Добавляем данные пользователя в контекст
		ctx := utils.SetUserContext(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RoleMiddleware проверяет роль пользователя, переданного в контексте
func RoleMiddleware(allowedRoles ...int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, role, err := utils.GetUserFromContext(r.Context())
			if err != nil {
				utils.ResponseWithError(w, http.StatusForbidden, "User not found in context")
				return
			}

			allowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					allowed = true
					break
				}
			}

			if !allowed {
				utils.ResponseWithError(w, http.StatusForbidden, "You do not have permission to access this resource")
				return
			}

			// Опционально: можно использовать userID для дополнительной логики

			next.ServeHTTP(w, r)
		})
	}
}
