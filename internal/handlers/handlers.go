package handlers

import (
	"encoding/json"
	"github.com/dqtu39/go-simple-server/internal/models"
	"github.com/dqtu39/go-simple-server/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (handler *EmployeeHandler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var employee *models.Employee
	employee = handler.service.GetEmployeeByID(id)
	if employee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(employee)

}

func (handler *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	handler.service.AddEmployee(employee)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

func (handler *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedEmployee models.Employee
	json.NewDecoder(r.Body).Decode(&updatedEmployee)
	success := handler.service.UpdateEmployee(id, updatedEmployee)
	if !success {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(updatedEmployee)
}

func (handler *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	success := handler.service.DeleteEmployee(id)
	if !success {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Employee deleted")
}

func (handler *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := handler.service.GetAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting all employees")
	}
	json.NewEncoder(w).Encode(employees)
}
