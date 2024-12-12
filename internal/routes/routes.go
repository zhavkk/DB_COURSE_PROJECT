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

	getClientHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 3)(
			http.HandlerFunc(controllers.GetClientHandler),
		),
	)
	r.Handle("/client/{id}", getClientHandler).Methods("GET", "OPTIONS")

	UpdateClientHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 3)(
			http.HandlerFunc(controllers.UpdateClientHandler),
		),
	)
	r.Handle("/client/{id}", UpdateClientHandler).Methods("PUT", "OPTIONS")
	GetServicesHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2, 3)(
			http.HandlerFunc(controllers.GetServicesHandler),
		),
	)
	r.Handle("/services", GetServicesHandler).Methods("GET", "OPTIONS")

	createServiceRequestHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2, 3)(
			http.HandlerFunc(controllers.CreateServiceRequestHandler),
		),
	)
	r.Handle("/create_service_requests", createServiceRequestHandler).Methods("POST", "OPTIONS")

	GetAllServiceRequests := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetServiceRequestsHandler),
		),
	)
	r.Handle("/service_requests", GetAllServiceRequests).Methods("GET", "OPTIONS")

	createServiceRequestEmployeeHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.CreateServiceRequestEmployeeHandler),
		),
	)
	r.Handle("/create_service_request_employees", createServiceRequestEmployeeHandler).Methods("POST", "OPTIONS")

	GetServiceRequestsForEmployeeId := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.GetServiceRequestsForEmployeeIdHandler),
		),
	)
	r.Handle("/service_request_employees", GetServiceRequestsForEmployeeId).Methods("GET", "OPTIONS")
	GetServiceRequestsForAdminsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetServiceRequestsForAdminsHandler),
		),
	)
	r.Handle("/service_request_employees_for_admins", GetServiceRequestsForAdminsHandler).Methods("GET", "OPTIONS")
	UpdateServiceRequestStatusHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.UpdateServiceRequestStatusHandler),
		),
	)
	r.Handle("/complete_service_request/{request_id}", UpdateServiceRequestStatusHandler).Methods("POST", "OPTIONS")

	// /service_reports - администратор и сотрудник
	CreateServiceReportHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1, 2)(
			http.HandlerFunc(controllers.CreateServiceReportHandler),
		),
	)
	r.Handle("/create_service_report", CreateServiceReportHandler).Methods("POST", "OPTIONS")

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

	GetAllServiceReportsHandler := auth.TokenVerifyMiddleware(
		auth.RoleMiddleware(1)(
			http.HandlerFunc(controllers.GetAllServiceReportsHandler),
		),
	)
	r.Handle("/all_service_reports", GetAllServiceReportsHandler).Methods("GET", "OPTIONS")

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
