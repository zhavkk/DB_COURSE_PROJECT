package routes

import (
	"dbproject/internal/auth"
	"dbproject/internal/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Маршруты для регистрации и логина
	r.HandleFunc("/login", auth.LoginUser).Methods("POST")
	r.HandleFunc("/register", auth.RegisterUser).Methods("POST")

	// Защищённый маршрут для получения всех пользователей
	// Добавляем middleware для проверки токена
	usersRouter := r.PathPrefix("/users").Subrouter() // создаём подмаршрут
	usersRouter.Use(auth.TokenVerifyMiddleware)       // добавляем middleware на подмаршрут
	usersRouter.HandleFunc("", controllers.GetUsersHandler).Methods("GET")

	// Добавьте сюда другие маршруты с защитой, например:
	// employeesRouter.Use(auth.TokenVerifyMiddleware)
	// employeesRouter.HandleFunc("", controllers.GetEmployeesHandler).Methods("GET")
}
