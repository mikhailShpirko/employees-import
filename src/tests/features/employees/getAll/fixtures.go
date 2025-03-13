package tests

import (
	employees "employees-import/features/employees"
	common "employees-import/tests/features/employees"
	"errors"

	"github.com/google/uuid"
)

type MockGetAllEmployeeExistsRepository struct {
}

func (repository *MockGetAllEmployeeExistsRepository) GetAll() ([]employees.Employee, error) {
	return []employees.Employee{*common.ValidExistingEmployee(uuid.New()), *common.ValidExistingEmployee(uuid.New()), *common.ValidExistingEmployee(uuid.New())}, nil
}

type MockFailGetAllGetAllEmployeeRepository struct{}

func (repository *MockFailGetAllGetAllEmployeeRepository) GetAll() ([]employees.Employee, error) {
	return nil, errors.New("FailGetAll")
}
