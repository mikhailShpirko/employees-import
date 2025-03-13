package endpoints_employees_create

import (
	custom_types "employees-import/customTypes"

	"github.com/google/uuid"
)

type ErrorCreatingEmployee struct {
	ErrorCodes []string `json:"errorCodes"`
}

type CreateEmployee struct {
	PayrollNumber   string                `json:"payrolNumber"`
	Forenames       string                `json:"forenames"`
	Surname         string                `json:"surname"`
	DateOfBirth     custom_types.DateOnly `json:"dateOfBirth"`
	TelephoneNumber string                `json:"telephoneNumber"`
	MobileNumber    string                `json:"mobileNumber"`
	AddressLine1    string                `json:"addressLine1"`
	AddressLine2    string                `json:"addressLine2"`
	Postcode        string                `json:"postcode"`
	Email           string                `json:"email"`
	StartDate       custom_types.DateOnly `json:"startDate"`
}

type EmployeeCreated struct {
	Id uuid.UUID `json:"id"`
}
