package tests

import (
	employees "employees-import/features/employees"
	"errors"

	"github.com/google/uuid"
)

type MockUpdateEmployeeRepository struct {
	PayrollExists bool
	IdExists      bool
}

func (repository *MockUpdateEmployeeRepository) IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error) {
	return repository.PayrollExists, nil
}

func (repository *MockUpdateEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return repository.IdExists, nil
}

func (repository *MockUpdateEmployeeRepository) Update(employee employees.Employee) error {
	return nil
}

type MockFailPayrollNumberExistsUpdateEmployeeRepository struct{}

func (repository *MockFailPayrollNumberExistsUpdateEmployeeRepository) IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error) {
	return false, errors.New("FailPayrollNumberExists")
}

func (repository *MockFailPayrollNumberExistsUpdateEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return true, nil
}

func (repository *MockFailPayrollNumberExistsUpdateEmployeeRepository) Update(employee employees.Employee) error {
	return nil
}

type MockFailIdExistsUpdateEmployeeRepository struct{}

func (repository *MockFailIdExistsUpdateEmployeeRepository) IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error) {
	return false, nil
}

func (repository *MockFailIdExistsUpdateEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return false, errors.New("FailIdExists")
}

func (repository *MockFailIdExistsUpdateEmployeeRepository) Update(employee employees.Employee) error {
	return nil
}

type MockFailUpdateUpdateEmployeeRepository struct{}

func (repository *MockFailUpdateUpdateEmployeeRepository) IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error) {
	return false, nil
}

func (repository *MockFailUpdateUpdateEmployeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	return true, nil
}

func (repository *MockFailUpdateUpdateEmployeeRepository) Update(employee employees.Employee) error {
	return errors.New("FailUpdate")
}
