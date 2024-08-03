package repository

import (
	"database/sql"
	"github.com/dqtu39/go-simple-server/internal/models"
)

type EmployeeRepository interface {
	GetAll() ([]models.Employee, error)
	GetByID(id int) (*models.Employee, error)
	Add(employee models.Employee) (int64, error)
	Update(id int, employee models.Employee) (int64, error)
	Delete(id int) (int64, error)
}

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.db.Query("SELECT * FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees := []models.Employee{}
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.EmployeeName, &employee.EmployeeSalary, &employee.EmployeeAge, &employee.ProfileImage); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *employeeRepository) GetByID(id int) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.QueryRow("SELECT * FROM employee WHERE ID = ?", id).Scan(&employee.ID, &employee.EmployeeName, &employee.EmployeeSalary, &employee.EmployeeAge, &employee.ProfileImage)
	if err != nil {
		return nil, err
	}
	return &employee, nil

}

func (r *employeeRepository) Add(employee models.Employee) (int64, error) {
	res, err := r.db.Exec("INSERT INTO employee (employee_name, employee_salary, employee_age, profile_image) VALUES (?, ?, ?, ?)", employee.EmployeeName, employee.EmployeeSalary, employee.EmployeeAge, employee.ProfileImage)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *employeeRepository) Update(id int, employee models.Employee) (int64, error) {
	result, err := r.db.Exec("UPDATE employee SET employee_name = ?, employee_salary = ?, employee_age = ?, profile_image = ? WHERE id = ?", employee.EmployeeName, employee.EmployeeSalary, employee.EmployeeAge, employee.ProfileImage, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *employeeRepository) Delete(id int) (int64, error) {
	result, err := r.db.Exec("DELETE FROM employee WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
