package tests

import (
	employees "employees-import/features/employees"
	create_handler "employees-import/features/employees/create"
	common "employees-import/tests/features/employees"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Create_Handle_ValidData_Created(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockCreateEmployeeRepository{PayrollExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := create_handler.Handle(employeeData, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Create Employee Handler returned error %v`, err)
	}

	switch createResult := result.(type) {
	case create_handler.Created:
		if createResult.Id == uuid.Nil {
			t.Fatalf(`Id of created record is empty`)
		}
	case create_handler.PayrollNumberAlreadyExists:
		t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
	case create_handler.ValidationErrors:
		t.Fatalf(`Unexpected result ValidationErrors %v`, createResult.Errors)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Create_Handle_InvalidValidData_ValidationErrors(t *testing.T) {
	employeeData := employees.EmployeeData{}

	repository := MockCreateEmployeeRepository{PayrollExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := create_handler.Handle(&employeeData, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Create Employee Handler returned error %v`, err)
	}

	switch result.(type) {
	case create_handler.Created:
		t.Fatalf(`Unexpected result Created`)
	case create_handler.PayrollNumberAlreadyExists:
		t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
	case create_handler.ValidationErrors:
		return
	default:
		t.Fatalf("Unsupported result")
		return
	}
}

func Test_Employees_Create_Handle_ExistingPayroll_PayrollNumberAlreadyExists(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockCreateEmployeeRepository{PayrollExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := create_handler.Handle(employeeData, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Create Employee Handler returned error %v`, err)
	}

	switch createResult := result.(type) {
	case create_handler.Created:
		t.Fatalf(`Unexpected result Created`)
	case create_handler.PayrollNumberAlreadyExists:
		return
	case create_handler.ValidationErrors:
		t.Fatalf(`Unexpected result ValidationErrors %v`, createResult.Errors)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Create_Handle_IsPayrollNumberExistReturnsError_Error(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockFailPayrollNumberExistsCreateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := create_handler.Handle(employeeData, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Create Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailPayrollNumberExists" {
		t.Fatalf(`Create Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Create_Handle_CreateReturnsError_Error(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockFailCreateCreateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := create_handler.Handle(employeeData, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Create Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailCreate" {
		t.Fatalf(`Create Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Create_Handle_UnitOfWorkReturnsError_Error(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockCreateEmployeeRepository{PayrollExists: false}

	unitOfWork := common.MockFailUnitOfWork{}

	result, err := create_handler.Handle(employeeData, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Create Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Create Employee Handler returned different error %v`, err.Error())
	}
}
