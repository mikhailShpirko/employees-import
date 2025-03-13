package employees_getAll

import (
	employees "employees-import/features/employees"
)

type IGetAllEmployeeRepository interface {
	GetAll() ([]employees.Employee, error)
}
