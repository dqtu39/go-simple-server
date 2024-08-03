package models

type Employee struct {
	ID           int    `json:"id"`
	EmployeeName string `json:"employee_name"`
	EmployeeAge  int    `json:"employee_age"`
	ProfileImage string `json:"profile_image"`
}
