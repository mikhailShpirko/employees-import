package employees_update

import (
	common_types "employees-import/features/common"
	employees "employees-import/features/employees"

	"time"
)

func Handle(employee *employees.Employee,
	repository IUpdateEmployeeRepository,
	unitOfWork common_types.IUnitOfWork) (IUpdateEmployeeResult, error) {

	validationErrors := employees.Validate(&employee.EmployeeData, time.Now())

	if validationErrors != nil {
		return ValidationErrors{Errors: validationErrors}, nil
	}

	isIdExist, idExistsError := repository.IsIdExist(employee.Id)

	if idExistsError != nil {
		return nil, idExistsError
	}

	if !isIdExist {
		return EmployeeNotExists{}, nil
	}

	isPayrollNumberExist, payrollExistsError := repository.IsPayrollNumberExistExclusive(employee.PayrollNumber, employee.Id)

	if payrollExistsError != nil {
		return nil, payrollExistsError
	}

	if isPayrollNumberExist {
		return PayrollNumberAlreadyExists{}, nil
	}

	updateError := repository.Update(employee)

	if updateError != nil {
		return nil, updateError
	}

	saveChangesError := unitOfWork.SaveChanges()

	if saveChangesError != nil {
		return nil, saveChangesError
	}

	return Updated{}, nil
}
