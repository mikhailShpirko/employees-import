package employees_create

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type ICreateEmployeeResult interface{}

type BaseCreateEmployeeResult struct{}

type Created struct {
	BaseCreateEmployeeResult
	Id uuid.UUID
}

type ValidationErrors struct {
	BaseCreateEmployeeResult

	Errors []employees.EMPLOYEE_VALIDATION_ERROR
}

type PayrollNumberAlreadyExists struct {
	BaseCreateEmployeeResult
}

type ICreateEmployeeRepository interface {
	IsPayrollNumberExist(payrollNumber string) (bool, error)
	Create(employee employees.Employee) error
}
