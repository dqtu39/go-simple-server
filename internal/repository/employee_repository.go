package repository

import "github.com/dqtu39/go-simple-server/internal/models"

type EmployeeRepository interface {
	GetAll() ([]models.Employee, error)
	GetByID(id int) *models.Employee
	Add(employee models.Employee)
	Update(id int, employee models.Employee) bool
	Delete(id int) bool
}

type employeeRepository struct {
	employees *[]models.Employee
}

func NewEmployeeRepository(employees *[]models.Employee) EmployeeRepository {
	return &employeeRepository{employees: employees}
}

func (r *employeeRepository) GetAll() ([]models.Employee, error) {
	return *r.employees, nil
}

func (r *employeeRepository) GetByID(id int) *models.Employee {
	for _, employee := range *r.employees {
		if employee.ID == id {
			return &employee
		}
	}
	return nil
}

func (r *employeeRepository) Add(employee models.Employee) {
	*r.employees = append(*r.employees, employee)
}

func (r *employeeRepository) Update(id int, employee models.Employee) bool {
	for i, e := range *r.employees {
		if e.ID == id {
			(*r.employees)[i] = employee
			return true
		}
	}
	return false
}

func (r *employeeRepository) Delete(id int) bool {
	for i, e := range *r.employees {
		if e.ID == id {
			*r.employees = append((*r.employees)[:i], (*r.employees)[i+1:]...)
			return true
		}
	}
	return false
}
