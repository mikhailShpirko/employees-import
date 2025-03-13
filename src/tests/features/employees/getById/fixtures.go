package tests

import (
	employees "employees-import/features/employees"
	common "employees-import/tests/features/employees"
	"errors"

	"github.com/google/uuid"
)

type MockGetByIdEmployeeExistsRepository struct {
}

func (repository *MockGetByIdEmployeeExistsRepository) GetById(id uuid.UUID) (bool, *employees.Employee, error) {
	return true, common.ValidExistingEmployee(uuid.New()), nil
}

type MockGetByIdEmployeeNotExistsRepository struct {
}

func (repository *MockGetByIdEmployeeNotExistsRepository) GetById(id uuid.UUID) (bool, *employees.Employee, error) {
	return false, nil, nil
}

type MockFailGetByIdGetByIdEmployeeRepository struct{}

func (repository *MockFailGetByIdGetByIdEmployeeRepository) GetById(id uuid.UUID) (bool, *employees.Employee, error) {
	return true, nil, errors.New("FailGetById")
}
