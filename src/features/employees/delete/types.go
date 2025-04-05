package employees_delete

import (
	"github.com/google/uuid"
)

type IDeleteEmployeeResult interface{}

func Match[T any](result IDeleteEmployeeResult,
	deletedDelegate func(deleted Deleted) T,
	employeeNotExistsDelegate func(employeeNotExists EmployeeNotExists) T) T {

	switch deleteResult := result.(type) {
	case Deleted:
		return deletedDelegate(deleteResult)
	case EmployeeNotExists:
		return employeeNotExistsDelegate(deleteResult)
	default:
		panic("Unsupported delete result")
	}
}

type Deleted struct {
}

type EmployeeNotExists struct {
}

type IDeleteEmployeeRepository interface {
	IsIdExist(id uuid.UUID) (bool, error)
	Delete(id uuid.UUID) error
}
