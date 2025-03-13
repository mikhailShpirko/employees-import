package employees_create

import (
	common_types "employees-import/features/common"
	employees "employees-import/features/employees"

	"time"

	"github.com/google/uuid"
)

func Handle(employeeData employees.EmployeeData,
	repository ICreateEmployeeRepository,
	unitOfWork common_types.IUnitOfWork) (ICreateEmployeeResult, error) {

	validationErrors := employees.Validate(employeeData, time.Now())

	if validationErrors != nil {
		return ValidationErrors{Errors: validationErrors}, nil
	}

	isPayrollNumberExist, payrollExistsError := repository.IsPayrollNumberExist(employeeData.PayrollNumber)

	if payrollExistsError != nil {
		return nil, payrollExistsError
	}

	if isPayrollNumberExist {
		return PayrollNumberAlreadyExists{}, nil
	}

	var employee employees.Employee
	employee.Id = uuid.New()
	employee.EmployeeData = employeeData

	createError := repository.Create(employee)

	if createError != nil {
		return nil, createError
	}

	saveChangesError := unitOfWork.SaveChanges()

	if saveChangesError != nil {
		return nil, saveChangesError
	}

	return Created{Id: employee.Id}, nil
}
