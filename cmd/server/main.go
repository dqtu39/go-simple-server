package main

import (
	"database/sql"
	"github.com/dqtu39/go-simple-server/internal/handlers"
	"github.com/dqtu39/go-simple-server/internal/repository"
	routes2 "github.com/dqtu39/go-simple-server/internal/routes"
	"github.com/dqtu39/go-simple-server/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:P@ssword@/godatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()

	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	} else {
		log.Println("Successfully connected to the database.")
	}

	// Initialize repository and service
	repo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(repo)
	handler := handlers.NewEmployeeHandler(employeeService)

	// Setup routes
	routes := routes2.SetupRoutes(handler)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
