package employees_getAll

import (
	employees "employees-import/features/employees"
)

func Handle(repository IGetAllEmployeeRepository) ([]employees.Employee, error) {

	employees, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	return employees, nil
}
