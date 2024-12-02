package routes

import (
	"dbproject/internal/controllers"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	//users
	r.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")
	//clients
	r.HandleFunc("/clients", controllers.GetClientHandler).Methods("GET")
	r.HandleFunc("/clients/", controllers.GetClientsHandler).Methods("GET")
	r.HandleFunc("/clients/create", controllers.CreateClientHandler).Methods("POST")
	r.HandleFunc("/clients/update", controllers.UpdateClientHandler).Methods("PUT")
	r.HandleFunc("/clients/delete", controllers.DeleteClientHandler).Methods("DELETE")
	//services
	r.HandleFunc("/services", controllers.CreateServiceHandler).Methods("POST")
	r.HandleFunc("/services/{id}", controllers.GetServiceHandler).Methods("GET")
	r.HandleFunc("/services", controllers.GetServicesHandler).Methods("GET")
	r.HandleFunc("/services/{id}", controllers.UpdateServiceHandler).Methods("PUT")
	r.HandleFunc("/services/{id}", controllers.DeleteServiceHandler).Methods("DELETE")
	//serviceRequests
	r.HandleFunc("/service_request", controllers.CreateServiceRequestHandler).Methods("POST")
	r.HandleFunc("/service_requests/{id}", controllers.GetServiceRequestHandler).Methods("GET")
	r.HandleFunc("/service_requests", controllers.GetServiceRequestsHandler).Methods("GET")
	r.HandleFunc("/service_requests/{id}", controllers.UpdateServiceRequestHandler).Methods("PUT")
	r.HandleFunc("/service_requests/{id}", controllers.DeleteServiceRequestHandler).Methods("DELETE")
	//serviceReports
	r.HandleFunc("/service-reports", controllers.GetAllServiceReports).Methods("GET")
	r.HandleFunc("/service-reports/", controllers.GetServiceReportByID).Methods("GET") // Используем URL-параметр для ID
	r.HandleFunc("/service-reports/create", controllers.CreateServiceReport).Methods("POST")
	r.HandleFunc("/service-reports/update", controllers.UpdateServiceReport).Methods("PUT")
	r.HandleFunc("/service-reports/delete", controllers.DeleteServiceReport).Methods("DELETE")
	//employee
	r.HandleFunc("/employees", controllers.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/", controllers.GetEmployeeByID).Methods("GET") // Используем URL-параметр для ID
	r.HandleFunc("/employees/{id}", controllers.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", controllers.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", controllers.UpdateEmployee).Methods("DELETE")
	return r
}
