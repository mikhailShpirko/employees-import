package tests

import (
	employees "employees-import/features/employees"
	"errors"

	"github.com/google/uuid"
)

type MockImportEmployeeRepository struct {
	PayrollIdMap map[string]uuid.UUID
}

func (repository *MockImportEmployeeRepository) GetPayrollNumberToIdMap() (map[string]uuid.UUID, error) {
	return repository.PayrollIdMap, nil
}

func (repository *MockImportEmployeeRepository) Create(employee employees.Employee) error {
	return nil
}

func (repository *MockImportEmployeeRepository) Update(employee employees.Employee) error {
	return nil
}

type MockFailGetPayrollNumberToIdMapImportEmployeeRepository struct {
	PayrollIdMap map[string]uuid.UUID
}

func (repository *MockFailGetPayrollNumberToIdMapImportEmployeeRepository) GetPayrollNumberToIdMap() (map[string]uuid.UUID, error) {
	return nil, errors.New("FailGetPayrollNumberToIdMap")
}

func (repository *MockFailGetPayrollNumberToIdMapImportEmployeeRepository) Create(employee employees.Employee) error {
	return nil
}

func (repository *MockFailGetPayrollNumberToIdMapImportEmployeeRepository) Update(employee employees.Employee) error {
	return nil
}

type MockFailCreateImportEmployeeRepository struct {
	PayrollIdMap map[string]uuid.UUID
}

func (repository *MockFailCreateImportEmployeeRepository) GetPayrollNumberToIdMap() (map[string]uuid.UUID, error) {
	return repository.PayrollIdMap, nil
}

func (repository *MockFailCreateImportEmployeeRepository) Create(employee employees.Employee) error {
	return errors.New("FailCreate")
}

func (repository *MockFailCreateImportEmployeeRepository) Update(employee employees.Employee) error {
	return nil
}

type MockFailUpdateImportEmployeeRepository struct {
	PayrollIdMap map[string]uuid.UUID
}

func (repository *MockFailUpdateImportEmployeeRepository) GetPayrollNumberToIdMap() (map[string]uuid.UUID, error) {
	return repository.PayrollIdMap, nil
}

func (repository *MockFailUpdateImportEmployeeRepository) Create(employee employees.Employee) error {
	return nil
}

func (repository *MockFailUpdateImportEmployeeRepository) Update(employee employees.Employee) error {
	return errors.New("FailUpdate")
}
