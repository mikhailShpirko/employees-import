package tests

import (
	employees "employees-import/features/employees"
	"errors"
)

type MockCreateEmployeeRepository struct {
	PayrollExists bool
}

func (repository *MockCreateEmployeeRepository) IsPayrollNumberExist(payrollNumber string) (bool, error) {
	return repository.PayrollExists, nil
}

func (repository *MockCreateEmployeeRepository) Create(employee *employees.Employee) error {
	return nil
}

type MockFailPayrollNumberExistsCreateEmployeeRepository struct{}

func (repository *MockFailPayrollNumberExistsCreateEmployeeRepository) IsPayrollNumberExist(payrollNumber string) (bool, error) {
	return false, errors.New("FailPayrollNumberExists")
}

func (repository *MockFailPayrollNumberExistsCreateEmployeeRepository) Create(employee *employees.Employee) error {
	return nil
}

type MockFailCreateCreateEmployeeRepository struct{}

func (repository *MockFailCreateCreateEmployeeRepository) IsPayrollNumberExist(payrollNumber string) (bool, error) {
	return false, nil
}

func (repository *MockFailCreateCreateEmployeeRepository) Create(employee *employees.Employee) error {
	return errors.New("FailCreate")
}
