package tests

import (
	delete_handler "employees-import/features/employees/delete"
	common "employees-import/tests/features/employees"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Delete_Handle_ExistingEmployee_Deleted(t *testing.T) {
	repository := MockDeleteEmployeeRepository{EmployeeExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := delete_handler.Handle(uuid.New(), &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Delete Employee Handler returned error %v`, err)
		return
	}

	switch result.(type) {
	case delete_handler.Deleted:
		return
	case delete_handler.EmployeeNotExists:
		t.Fatalf(`Unexpected result EmployeeNotExists`)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Delete_Handle_IdNotExists_EmployeeNotExists(t *testing.T) {
	repository := MockDeleteEmployeeRepository{EmployeeExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := delete_handler.Handle(uuid.New(), &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Delete Employee Handler returned error %v`, err)
	}

	switch result.(type) {
	case delete_handler.Deleted:
		t.Fatalf(`Unexpected result Deleted`)
	case delete_handler.EmployeeNotExists:
		return
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Delete_Handle_IsIdExistReturnsError_Error(t *testing.T) {

	repository := MockFailIdExistsDeleteEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := delete_handler.Handle(uuid.New(), &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Delete Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailIdExists" {
		t.Fatalf(`Delete Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Delete_Handle_DeleteReturnsError_Error(t *testing.T) {
	repository := MockFailDeleteDeleteEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := delete_handler.Handle(uuid.New(), &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Delete Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailDelete" {
		t.Fatalf(`Delete Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Delete_Handle_UnitOfWorkReturnsError_Error(t *testing.T) {
	repository := MockDeleteEmployeeRepository{EmployeeExists: true}

	unitOfWork := common.MockFailUnitOfWork{}

	result, err := delete_handler.Handle(uuid.New(), &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Delete Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Delete Employee Handler returned different error %v`, err.Error())
	}
}
