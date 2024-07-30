package storage

import (
	"encoding/json"
	"github.com/dqtu39/go-simple-server/internal/models"
	"io/ioutil"
	"os"
)

var Employees []models.Employee

func LoadEmployees(filename string) error {
	file, err := os.Open(filename)
	if (err != nil) && os.IsNotExist(err) {
		return SaveEmployees(filename)
	} else if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Employees)
	if err != nil {
		return err
	}

	return nil
}

func SaveEmployees(filename string) error {
	data, err := json.MarshalIndent(Employees, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
