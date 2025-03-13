package tests

import (
	getAll_handler "employees-import/features/employees/getAll"
	"testing"
)

func Test_Employees_GetAll_Handle_Employees_AllEmployeesReturned(t *testing.T) {
	repository := MockGetAllEmployeeExistsRepository{}

	employees, err := getAll_handler.Handle(&repository)

	if err != nil {
		t.Fatalf(`GetAll Employee Handler returned error %v`, err)
		return
	}

	if employees == nil {
		t.Fatalf(`GetAll Employee Handler nil`)
		return
	}

	if len(employees) != 3 {
		t.Fatalf(`GetAll Employee Handler expected 3 employees but returned %v`, employees)
		return
	}
}

func Test_Employees_GetAll_Handle_Error_ErrorReturned(t *testing.T) {
	repository := MockFailGetAllGetAllEmployeeRepository{}

	result, err := getAll_handler.Handle(&repository)

	if err == nil {
		t.Fatalf(`GetAll Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailGetAll" {
		t.Fatalf(`GetAllEmployee Handler returned different error %v`, err.Error())
	}
}
