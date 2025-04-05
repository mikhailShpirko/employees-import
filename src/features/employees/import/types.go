package employees_import

import (
	employees "employees-import/features/employees"

	"github.com/google/uuid"
)

type EMPLOYEE_IMPORT_STATUS int

const (
	EMPLOYEE_IMPORT_STATUS_CREATED EMPLOYEE_IMPORT_STATUS = iota
	EMPLOYEE_IMPORT_STATUS_UPDATED
)

func (importStatus EMPLOYEE_IMPORT_STATUS) String() string {
	return [...]string{"CREATED",
		"UPDATED"}[importStatus]
}

type EmployeeImportResult struct {
	PayrollNumber string
	Status        EMPLOYEE_IMPORT_STATUS
	Id            uuid.UUID
}

type InvalidEmployeeData struct {
	PayrollNumber    string
	ValidationErrors []employees.EMPLOYEE_VALIDATION_ERROR
}

type IImportEmployeesResult interface{}

func Match[T any](result IImportEmployeesResult,
	importedDelegate func(imported SuccessfullyImported) T,
	validationErrorsDelegate func(validationErrors ValidationErrors) T) T {

	switch importResult := result.(type) {
	case SuccessfullyImported:
		return importedDelegate(importResult)
	case ValidationErrors:
		return validationErrorsDelegate(importResult)
	default:
		panic("Unsupported import result")
	}
}

type SuccessfullyImported struct {
	Result []EmployeeImportResult
}

type ValidationErrors struct {
	Errors []InvalidEmployeeData
}

type IImportEmployeeRepository interface {
	GetPayrollNumberToIdMap() (map[string]uuid.UUID, error)
	Create(employee employees.Employee) error
	Update(employee employees.Employee) error
}
