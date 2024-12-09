package routes

import (
	"dbproject/internal/auth"
	"dbproject/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes - настройка всех маршрутов
func SetupRoutes(r *mux.Router) {
	// Добавляем middleware для CORS
	r.Use(auth.CORS)

	// Маршруты для аутентификации и регистрации
	r.HandleFunc("/login", auth.LoginUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/register", auth.RegisterUser).Methods("POST", "OPTIONS")

	// Защищённые маршруты для пользователей, доступные только после аутентификации

	// /users - только администратор
	usersHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetUsersHandler),
		),
	)
	r.Handle("/users", usersHandler).Methods("GET", "OPTIONS")

	// /clients - администратор и сотрудник
	clientsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetClientsHandler),
		),
	)
	r.Handle("/clients", clientsHandler).Methods("GET", "OPTIONS")
	GetClientByUSERIDHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2, 3)(
			http.HandlerFunc(controllers.GetClientByUSERIDHandler),
		),
	)
	r.Handle("/getClientId", GetClientByUSERIDHandler).Methods("GET", "OPTIONS")

	GetEmployeeByUSERIDHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2, 3)(
			http.HandlerFunc(controllers.GetEmployeeByUSERIDHandler),
		),
	)
	r.Handle("/getEmployeeId", GetEmployeeByUSERIDHandler).Methods("GET", "OPTIONS")

	GetServicesHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(2, 3)(
			http.HandlerFunc(controllers.GetServicesHandler),
		),
	)
	r.Handle("/services", GetServicesHandler).Methods("GET", "OPTIONS")

	createServiceRequestHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(2, 3)(
			http.HandlerFunc(controllers.CreateServiceRequestHandler),
		),
	)
	r.Handle("/service_requests", createServiceRequestHandler).Methods("POST", "OPTIONS")

	// /service_reports - администратор и сотрудник
	serviceReportsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetAllServiceReportsHandler),
		),
	)
	r.Handle("/service_reports", serviceReportsHandler).Methods("GET", "OPTIONS")

	// /service_requests/{id} - администратор и сотрудник
	serviceRequestByIDHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetServiceRequestHandler),
		),
	)
	r.Handle("/service_requests/{id:[0-9]+}", serviceRequestByIDHandler).Methods("GET", "OPTIONS")

	// Создание отчёта - сотрудник и администратор
	createServiceReportHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.CreateServiceReportHandler),
		),
	)
	r.Handle("/service_reports", createServiceReportHandler).Methods("POST", "OPTIONS")

	// /employees - только администратор
	employeesHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetEmployeesHandler),
		),
	)
	r.Handle("/employees", employeesHandler).Methods("GET", "OPTIONS")

	// /roles - только администратор
	rolesHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetRolesHandler),
		),
	)
	r.Handle("/roles", rolesHandler).Methods("GET", "OPTIONS")
}
