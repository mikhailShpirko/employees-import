package employees_delete

import (
	common_types "employees-import/features/common"

	"github.com/google/uuid"
)

func Handle(employeeId uuid.UUID,
	repository IDeleteEmployeeRepository,
	unitOfWork common_types.IUnitOfWork) (IDeleteEmployeeResult, error) {

	isEmployeeExist, payrollExistsError := repository.IsIdExist(employeeId)

	if payrollExistsError != nil {
		return nil, payrollExistsError
	}

	if !isEmployeeExist {
		return EmployeeNotExists{}, nil
	}

	deleteError := repository.Delete(employeeId)

	if deleteError != nil {
		return nil, deleteError
	}

	saveChangesError := unitOfWork.SaveChanges()

	if saveChangesError != nil {
		return nil, saveChangesError
	}

	return Deleted{}, nil
}
