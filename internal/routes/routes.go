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
	r.HandleFunc("/login", auth.LoginUser).Methods("POST")
	r.HandleFunc("/register", auth.RegisterUser).Methods("POST")

	// Защищённые маршруты для пользователей, доступные только после аутентификации

	// /users - только администратор
	usersHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetUsersHandler),
		),
	)
	r.Handle("/users", usersHandler).Methods("GET")

	// /clients - администратор и сотрудник
	clientsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetClientsHandler),
		),
	)
	r.Handle("/clients", clientsHandler).Methods("GET")

	// /service_requests - администратор и сотрудник
	serviceRequestsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetServiceRequestsHandler),
		),
	)
	r.Handle("/service_requests", serviceRequestsHandler).Methods("GET")

	// /service_reports - администратор и сотрудник
	serviceReportsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetAllServiceReportsHandler),
		),
	)
	r.Handle("/service_reports", serviceReportsHandler).Methods("GET")

	// /service_requests/{id} - администратор и сотрудник
	serviceRequestByIDHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetServiceRequestHandler),
		),
	)
	r.Handle("/service_requests/{id:[0-9]+}", serviceRequestByIDHandler).Methods("GET")

	// Создание заявки - клиент и сотрудник
	createServiceRequestHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(2, 3)(
			http.HandlerFunc(controllers.CreateServiceRequestHandler),
		),
	)
	r.Handle("/service_requests", createServiceRequestHandler).Methods("POST")

	// Создание отчёта - сотрудник и администратор
	createServiceReportHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.CreateServiceReportHandler),
		),
	)
	r.Handle("/service_reports", createServiceReportHandler).Methods("POST")

	// /employees - только администратор
	employeesHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetEmployeesHandler),
		),
	)
	r.Handle("/employees", employeesHandler).Methods("GET")

	// /roles - только администратор
	rolesHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetRolesHandler),
		),
	)
	r.Handle("/roles", rolesHandler).Methods("GET")
}
