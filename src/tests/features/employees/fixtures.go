package tests

import (
	employees "employees-import/features/employees"
	"errors"
	"time"

	"github.com/google/uuid"
)

func ValidNewEmployee() *employees.EmployeeData {
	dateOfBirth, _ := time.Parse(time.DateOnly, "1990-12-11")
	startDate, _ := time.Parse(time.DateOnly, "2025-10-24")

	employee := employees.CreateEmployeeData("TEST_PayrollNumber",
		"TEST_Forenames",
		"TEST_Surname",
		dateOfBirth,
		"998901234567",
		"998129834567",
		"TEST_AddressLine1",
		"TEST_AddressLine2",
		"TEST_Postcode",
		"test@email.com",
		startDate)

	return employee
}

func ValidExistingEmployee(id uuid.UUID) *employees.Employee {
	employeeData := ValidNewEmployee()

	return &employees.Employee{Id: id, EmployeeData: *employeeData}
}

type MockSuccessUnitOfWork struct{}

func (unitOfWork *MockSuccessUnitOfWork) SaveChanges() error {
	return nil
}

type MockFailUnitOfWork struct{}

func (unitOfWork *MockFailUnitOfWork) SaveChanges() error {
	return errors.New("FailUnitOfWork")
}
