package tests

import (
	employees_getById "employees-import/features/employees/getById"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_GetById_Handle_ExistingEmployee_EmployeeExists(t *testing.T) {
	repository := MockGetByIdEmployeeExistsRepository{}

	result, err := employees_getById.Handle(uuid.New(), &repository)

	if err != nil {
		t.Fatalf(`GetById Employee Handler returned error %v`, err)
	}

	employees_getById.Match(result,
		func(exists employees_getById.EmployeeExists) error {
			return nil
		},
		func(notExists employees_getById.EmployeeNotExists) error {
			t.Fatalf(`Unexpected result EmployeeExists`)
			return nil
		})
}

func Test_Employees_GetById_Handle_NotExistingEmployee_EmployeeNotExists(t *testing.T) {
	repository := MockGetByIdEmployeeNotExistsRepository{}

	result, err := employees_getById.Handle(uuid.New(), &repository)

	if err != nil {
		t.Fatalf(`GetById Employee Handler returned error %v`, err)
	}

	employees_getById.Match(result,
		func(exists employees_getById.EmployeeExists) error {
			t.Fatalf(`Unexpected result EmployeeExists`)
			return nil
		},
		func(notExists employees_getById.EmployeeNotExists) error {
			return nil
		})
}

func Test_Employees_GetById_Handle_Error_Error(t *testing.T) {

	repository := MockFailGetByIdGetByIdEmployeeRepository{}

	result, err := employees_getById.Handle(uuid.New(), &repository)

	if err == nil {
		t.Fatalf(`GetById Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailGetById" {
		t.Fatalf(`GetById Employee Handler returned different error %v`, err.Error())
	}
}
