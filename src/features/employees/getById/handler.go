package employees_getById

import (
	"github.com/google/uuid"
)

func Handle(employeeId uuid.UUID,
	repository IGetByIdEmployeeRepository) (IGetByIdEmployeeResult, error) {

	isExists, employee, err := repository.GetById(employeeId)

	if err != nil {
		return nil, err
	}

	if !isExists {
		return EmployeeNotExists{}, nil
	}

	return EmployeeExists{Employee: employee}, nil
}
