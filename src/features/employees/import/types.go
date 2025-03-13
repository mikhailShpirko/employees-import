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

type BaseImportEmployeesResult struct{}

type SuccessfullyImported struct {
	BaseImportEmployeesResult
	Result []EmployeeImportResult
}

type ValidationErrors struct {
	BaseImportEmployeesResult
	Errors []InvalidEmployeeData
}

type IImportEmployeeRepository interface {
	GetPayrollNumberToIdMap() (map[string]uuid.UUID, error)
	Create(employee employees.Employee) error
	Update(employee employees.Employee) error
}
