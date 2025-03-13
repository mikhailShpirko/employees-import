package tests

import (
	"errors"

	"github.com/google/uuid"
)

type MockDeleteEmployeeRepository struct {
	EmployeeExists bool
}

func (repository *MockDeleteEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return repository.EmployeeExists, nil
}

func (repository *MockDeleteEmployeeRepository) Delete(id uuid.UUID) error {
	return nil
}

type MockFailIdExistsDeleteEmployeeRepository struct{}

func (repository *MockFailIdExistsDeleteEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return false, errors.New("FailIdExists")
}

func (repository *MockFailIdExistsDeleteEmployeeRepository) Delete(id uuid.UUID) error {
	return nil
}

type MockFailDeleteDeleteEmployeeRepository struct{}

func (repository *MockFailDeleteDeleteEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return true, nil
}

func (repository *MockFailDeleteDeleteEmployeeRepository) Delete(id uuid.UUID) error {
	return errors.New("FailDelete")
}
