package employees_update

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type IUpdateEmployeeResult interface{}

func Match[T any](result IUpdateEmployeeResult,
	updatedDelegate func(updated Updated) T,
	payrollExistsDelegate func(payrollExists PayrollNumberAlreadyExists) T,
	validationErrorsDelegate func(validationErrors ValidationErrors) T,
	notExistsDelegate func(notExists EmployeeNotExists) T) T {

	switch updateResult := result.(type) {
	case Updated:
		return updatedDelegate(updateResult)
	case PayrollNumberAlreadyExists:
		return payrollExistsDelegate(updateResult)
	case ValidationErrors:
		return validationErrorsDelegate(updateResult)
	case EmployeeNotExists:
		return notExistsDelegate(updateResult)
	default:
		panic("Unsupported update result")
	}
}

type Updated struct {
}

type ValidationErrors struct {
	Errors []employees.EMPLOYEE_VALIDATION_ERROR
}

type PayrollNumberAlreadyExists struct {
}

type EmployeeNotExists struct {
}

type IUpdateEmployeeRepository interface {
	IsIdExist(id uuid.UUID) (bool, error)
	IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error)
	Update(employee employees.Employee) error
}
