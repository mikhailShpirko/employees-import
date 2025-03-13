package employees_getById

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type IGetByIdEmployeeResult interface{}

type BaseGetByIdEmployeeResult struct{}

type EmployeeExists struct {
	BaseGetByIdEmployeeResult

	Employee *employees.Employee
}

type EmployeeNotExists struct {
	BaseGetByIdEmployeeResult
}

type IGetByIdEmployeeRepository interface {
	GetById(id uuid.UUID) (bool, *employees.Employee, error)
}
