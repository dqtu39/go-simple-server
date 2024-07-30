package routes

import (
	"github.com/dqtu39/go-simple-server/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(handlers *handlers.EmployeeHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/employees", handlers.GetAllEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", handlers.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employees", handlers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", handlers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", handlers.DeleteEmployee).Methods("DELETE")
	return router
}
