package employees_create

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type ICreateEmployeeResult interface{}

func Match[T any](result ICreateEmployeeResult,
	createdDelegate func(created Created) T,
	payrollExistsDelegate func(payrollExists PayrollNumberAlreadyExists) T,
	validationErrorsDelegate func(validationErrors ValidationErrors) T) T {

	switch createResult := result.(type) {
	case Created:
		return createdDelegate(createResult)
	case PayrollNumberAlreadyExists:
		return payrollExistsDelegate(createResult)
	case ValidationErrors:
		return validationErrorsDelegate(createResult)
	default:
		panic("Unsupported create result")
	}
}

type Created struct {
	Id uuid.UUID
}

type ValidationErrors struct {
	Errors []employees.EMPLOYEE_VALIDATION_ERROR
}

type PayrollNumberAlreadyExists struct {
}

type ICreateEmployeeRepository interface {
	IsPayrollNumberExist(payrollNumber string) (bool, error)
	Create(employee employees.Employee) error
}
