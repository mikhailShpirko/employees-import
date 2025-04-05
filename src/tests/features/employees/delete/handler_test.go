package tests

import (
	employees_delete "employees-import/features/employees/delete"
	common "employees-import/tests/features/employees"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Delete_Handle_ExistingEmployee_Deleted(t *testing.T) {
	repository := MockDeleteEmployeeRepository{EmployeeExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_delete.Handle(uuid.New(), &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Delete Employee Handler returned error %v`, err)
		return
	}

	employees_delete.Match(result,
		func(_ employees_delete.Deleted) error {
			return nil
		},
		func(employeeNotExists employees_delete.EmployeeNotExists) error {
			t.Fatalf(`Unexpected result EmployeeNotExists`)
			return nil
		})
}

func Test_Employees_Delete_Handle_IdNotExists_EmployeeNotExists(t *testing.T) {
	repository := MockDeleteEmployeeRepository{EmployeeExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_delete.Handle(uuid.New(), &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Delete Employee Handler returned error %v`, err)
	}

	employees_delete.Match(result,
		func(_ employees_delete.Deleted) error {
			t.Fatalf(`Unexpected result Deleted`)
			return nil
		},
		func(employeeNotExists employees_delete.EmployeeNotExists) error {
			return nil
		})
}

func Test_Employees_Delete_Handle_IsIdExistReturnsError_Error(t *testing.T) {

	repository := MockFailIdExistsDeleteEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_delete.Handle(uuid.New(), &repository, &unitOfWork)

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

	result, err := employees_delete.Handle(uuid.New(), &repository, &unitOfWork)

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

	result, err := employees_delete.Handle(uuid.New(), &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Delete Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Delete Employee Handler returned different error %v`, err.Error())
	}
}
