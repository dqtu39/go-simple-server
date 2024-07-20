package main

import (
	"net/http"
)

type Employee struct {
	Id           int    `json:"id"`
	Name         string `json:"employee_name"`
	Age          int    `json:"employee_age"`
	ProfileImage string `json:"profile_image"`
}

func main() {
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
