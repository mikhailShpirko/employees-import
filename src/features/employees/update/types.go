package employees_update

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type IUpdateEmployeeResult interface{}

type BaseUpdateEmployeeResult struct{}

type Updated struct {
	BaseUpdateEmployeeResult
}

type ValidationErrors struct {
	BaseUpdateEmployeeResult

	Errors []employees.EMPLOYEE_VALIDATION_ERROR
}

type PayrollNumberAlreadyExists struct {
	BaseUpdateEmployeeResult
}

type EmployeeNotExists struct {
	BaseUpdateEmployeeResult
}

type IUpdateEmployeeRepository interface {
	IsIdExist(id uuid.UUID) (bool, error)
	IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error)
	Update(employee employees.Employee) error
}
