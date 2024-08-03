package main

import (
	"github.com/dqtu39/go-simple-server/internal/handlers"
	"github.com/dqtu39/go-simple-server/internal/repository"
	"github.com/dqtu39/go-simple-server/internal/routes"
	"github.com/dqtu39/go-simple-server/internal/service"
	"github.com/dqtu39/go-simple-server/internal/storage"
	"log"
	"net/http"
)

func main() {
	err := storage.LoadEmployees("data/employees.json")
	if err != nil {
		log.Fatalf("Failed to load employees: %v", err)
	}

	repo := repository.NewEmployeeRepository(&storage.Employees)
	employeeService := service.NewEmployeeService(repo)
	handler := handlers.NewEmployeeHandler(employeeService)

	router := routes.SetupRoutes(handler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
