package tests

import (
	employees "employees-import/features/employees"
	"slices"
	"testing"
	"time"
)

func Test_ValidateEmployee_ValidData_NoErrors(t *testing.T) {
	employee := ValidNewEmployee()

	today, _ := time.Parse(time.DateOnly, "2024-02-16")

	errors := employees.Validate(employee, today)
	if errors != nil {
		t.Fatalf(`Validate returned following errors %v`, errors)
	}
}

func Test_ValidateEmployee_EmptyRequiredFields_ErrorsReturned(t *testing.T) {
	employee := employees.EmployeeData{}

	today, _ := time.Parse(time.DateOnly, "2024-02-16")

	errors := employees.Validate(employee, today)
	if errors == nil {
		t.Fatalf(`Validate didn't return any errors`)
	}

	if len(errors) != 8 {
		t.Fatalf(`There should be 8 validation errors, but in fact are %v`, len(errors))
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_PAYROLL_NUMBER) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_PAYROLL_NUMBER missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_FORENAMES) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_FORENAMES missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_SURNAME) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_SURNAME missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_DATE_OF_BIRTH) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_DATE_OF_BIRTH missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_TELEPHONE_NUMBER) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_TELEPHONE_NUMBER missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_MOBILE_NUMBER) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_MOBILE_NUMBER missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_EMAIL) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_EMAIL missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_EMPTY_START_DATE) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_EMPTY_START_DATE missing`)
	}
}

func Test_ValidateEmployee_InvalidFields_ErrorsReturned(t *testing.T) {
	employee := ValidNewEmployee()

	futureDateOfBirth, _ := time.Parse(time.DateOnly, "2024-03-16")
	employee.DateOfBirth = futureDateOfBirth
	employee.TelephoneNumber = "ABCDEFG"
	employee.MobileNumber = "ABCDEFG"
	employee.Email = "ABCDEFG"

	today, _ := time.Parse(time.DateOnly, "2024-02-16")

	errors := employees.Validate(employee, today)
	if errors == nil {
		t.Fatalf(`Validate didn't return any errors`)
	}

	if len(errors) != 4 {
		t.Fatalf(`There should be 4 validation errors, but in fact are %v`, len(errors))
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_FUTURE_DATE_OF_BIRTH) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_FUTURE_DATE_OF_BIRTH missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_INVALID_TELEPHONE_NUMBER) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_INVALID_TELEPHONE_NUMBER missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_INVALID_MOBILE_NUMBER) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_INVALID_MOBILE_NUMBER missing`)
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_INVALID_EMAIL) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_INVALID_EMAIL missing`)
	}
}

func Test_ValidateEmployee_InvalidDateFieldsFields_ErrorsReturned(t *testing.T) {
	employee := ValidNewEmployee()

	newDateOfBirth, _ := time.Parse(time.DateOnly, "2022-03-16")
	newStartDate, _ := time.Parse(time.DateOnly, "2022-02-16")
	employee.DateOfBirth = newDateOfBirth
	employee.StartDate = newStartDate

	today, _ := time.Parse(time.DateOnly, "2024-02-16")

	errors := employees.Validate(employee, today)
	if errors == nil {
		t.Fatalf(`Validate didn't return any errors`)
	}

	if len(errors) != 1 {
		t.Fatalf(`There should be 1 validation errors, but in fact are %v`, len(errors))
	}

	if !slices.Contains(errors, employees.EMPLOYEE_VALIDATION_ERROR_DATE_OF_BIRTH_AFTER_START_DATE) {
		t.Fatalf(`EMPLOYEE_VALIDATION_ERROR_DATE_OF_BIRTH_AFTER_START_DATE missing`)
	}
}
