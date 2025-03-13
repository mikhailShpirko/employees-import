package employees_delete

import (
	"github.com/google/uuid"
)

type IDeleteEmployeeResult interface{}

type BaseDeleteEmployeeResult struct{}

type Deleted struct {
	BaseDeleteEmployeeResult
}

type EmployeeNotExists struct {
	BaseDeleteEmployeeResult
}

type IDeleteEmployeeRepository interface {
	IsIdExist(id uuid.UUID) (bool, error)
	Delete(id uuid.UUID) error
}
