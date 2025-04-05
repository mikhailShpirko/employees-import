package tests

import (
	employees "employees-import/features/employees"
	employees_create "employees-import/features/employees/create"
	common "employees-import/tests/features/employees"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Create_Handle_ValidData_Created(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockCreateEmployeeRepository{PayrollExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_create.Handle(employeeData, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Create Employee Handler returned error %v`, err)
	}

	employees_create.Match(result,
		func(created employees_create.Created) error {
			if created.Id == uuid.Nil {
				t.Fatalf(`Id of created record is empty`)
			}
			return nil
		},
		func(payrollExists employees_create.PayrollNumberAlreadyExists) error {
			t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
			return nil
		},
		func(validationErrors employees_create.ValidationErrors) error {
			t.Fatalf(`Unexpected result ValidationErrors %v`, validationErrors.Errors)
			return nil
		})
}

func Test_Employees_Create_Handle_InvalidValidData_ValidationErrors(t *testing.T) {
	employeeData := employees.EmployeeData{}

	repository := MockCreateEmployeeRepository{PayrollExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_create.Handle(employeeData, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Create Employee Handler returned error %v`, err)
	}

	employees_create.Match(result,
		func(created employees_create.Created) error {
			t.Fatalf(`Unexpected result Created`)
			return nil
		},
		func(payrollExists employees_create.PayrollNumberAlreadyExists) error {
			t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
			return nil
		},
		func(validationErrors employees_create.ValidationErrors) error {
			return nil
		})
}

func Test_Employees_Create_Handle_ExistingPayroll_PayrollNumberAlreadyExists(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockCreateEmployeeRepository{PayrollExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_create.Handle(employeeData, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Create Employee Handler returned error %v`, err)
	}

	employees_create.Match(result,
		func(created employees_create.Created) error {
			t.Fatalf(`Unexpected result Created`)
			return nil
		},
		func(payrollExists employees_create.PayrollNumberAlreadyExists) error {
			return nil
		},
		func(validationErrors employees_create.ValidationErrors) error {
			t.Fatalf(`Unexpected result ValidationErrors %v`, validationErrors.Errors)
			return nil
		})
}

func Test_Employees_Create_Handle_IsPayrollNumberExistReturnsError_Error(t *testing.T) {
	employeeData := common.ValidNewEmployee()

	repository := MockFailPayrollNumberExistsCreateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_create.Handle(employeeData, &repository, &unitOfWork)

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

	result, err := employees_create.Handle(employeeData, &repository, &unitOfWork)

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

	result, err := employees_create.Handle(employeeData, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Create Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Create Employee Handler returned different error %v`, err.Error())
	}
}
