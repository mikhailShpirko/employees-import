package tests

import (
	getById_handler "employees-import/features/employees/getById"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_GetById_Handle_ExistingEmployee_EmployeeExists(t *testing.T) {
	repository := MockGetByIdEmployeeExistsRepository{}

	result, err := getById_handler.Handle(uuid.New(), &repository)

	if err != nil {
		t.Fatalf(`GetById Employee Handler returned error %v`, err)
	}

	switch result.(type) {
	case getById_handler.EmployeeExists:
		return
	case getById_handler.EmployeeNotExists:
		t.Fatalf(`Unexpected result EmployeeNotExists`)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_GetById_Handle_NotExistingEmployee_EmployeeNotExists(t *testing.T) {
	repository := MockGetByIdEmployeeNotExistsRepository{}

	result, err := getById_handler.Handle(uuid.New(), &repository)

	if err != nil {
		t.Fatalf(`GetById Employee Handler returned error %v`, err)
	}

	switch result.(type) {
	case getById_handler.EmployeeExists:
		t.Fatalf(`Unexpected result EmployeeExists`)
	case getById_handler.EmployeeNotExists:
		return
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_GetById_Handle_Error_Error(t *testing.T) {

	repository := MockFailGetByIdGetByIdEmployeeRepository{}

	result, err := getById_handler.Handle(uuid.New(), &repository)

	if err == nil {
		t.Fatalf(`GetById Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailGetById" {
		t.Fatalf(`GetById Employee Handler returned different error %v`, err.Error())
	}
}
