package employees

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeData struct {
	PayrollNumber   string
	Forenames       string
	Surname         string
	DateOfBirth     time.Time
	TelephoneNumber string
	MobileNumber    string
	AddressLine1    string
	AddressLine2    string
	Postcode        string
	Email           string
	StartDate       time.Time
}

type Employee struct {
	EmployeeData
	Id uuid.UUID
}

func CreateEmployeeData(
	payrollNumber string,
	forenames string,
	surname string,
	dateOfBirth time.Time,
	telephoneNumber string,
	mobileNumber string,
	addressLine1 string,
	addressLine2 string,
	postcode string,
	email string,
	startDate time.Time) EmployeeData {

	var employee EmployeeData
	employee.PayrollNumber = payrollNumber
	employee.Forenames = forenames
	employee.Surname = surname
	employee.DateOfBirth = dateOfBirth
	employee.TelephoneNumber = telephoneNumber
	employee.MobileNumber = mobileNumber
	employee.AddressLine1 = addressLine1
	employee.AddressLine2 = addressLine2
	employee.Postcode = postcode
	employee.Email = email
	employee.StartDate = startDate

	return employee
}

func CreateEmployee(
	id uuid.UUID,
	payrollNumber string,
	forenames string,
	surname string,
	dateOfBirth time.Time,
	telephoneNumber string,
	mobileNumber string,
	addressLine1 string,
	addressLine2 string,
	postcode string,
	email string,
	startDate time.Time) Employee {

	var employee Employee
	employee.EmployeeData = CreateEmployeeData(payrollNumber,
		forenames,
		surname,
		dateOfBirth,
		telephoneNumber,
		mobileNumber,
		addressLine1,
		addressLine2,
		postcode,
		email,
		startDate)

	employee.Id = id

	return employee
}
