package employees_getById

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type IGetByIdEmployeeResult interface{}

func Match[T any](result IGetByIdEmployeeResult,
	existsDelegate func(exists EmployeeExists) T,
	notExistsDelegate func(notExists EmployeeNotExists) T) T {

	switch getByIdResult := result.(type) {
	case EmployeeExists:
		return existsDelegate(getByIdResult)
	case EmployeeNotExists:
		return notExistsDelegate(getByIdResult)
	default:
		panic("Unsupported get by id result")
	}
}

type EmployeeExists struct {
	Employee employees.Employee
}

type EmployeeNotExists struct {
}

type IGetByIdEmployeeRepository interface {
	GetById(id uuid.UUID) (bool, employees.Employee, error)
}
