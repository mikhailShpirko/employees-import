package tests

import (
	employees "employees-import/features/employees"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_NewEmployee_FieldsMappedProperly(t *testing.T) {

	dateOfBirth, _ := time.Parse(time.DateOnly, "1990-12-11")
	startDate, _ := time.Parse(time.DateOnly, "2025-10-24")

	employee := employees.NewEmployee("TEST_PayrollNumber",
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

	if employee.PayrollNumber != "TEST_PayrollNumber" {
		t.Fatalf(`PayrollNumber mapped improperly`)
	}

	if employee.Forenames != "TEST_Forenames" {
		t.Fatalf(`Forenames mapped improperly`)
	}

	if employee.Surname != "TEST_Surname" {
		t.Fatalf(`Surname mapped improperly`)
	}

	if employee.DateOfBirth != dateOfBirth {
		t.Fatalf(`DateOfBirth mapped improperly`)
	}

	if employee.TelephoneNumber != "998901234567" {
		t.Fatalf(`TelephoneNumber mapped improperly`)
	}

	if employee.MobileNumber != "998129834567" {
		t.Fatalf(`MobileNumber mapped improperly`)
	}

	if employee.AddressLine1 != "TEST_AddressLine1" {
		t.Fatalf(`AddressLine1 mapped improperly`)
	}

	if employee.AddressLine2 != "TEST_AddressLine2" {
		t.Fatalf(`AddressLine2 mapped improperly`)
	}

	if employee.Postcode != "TEST_Postcode" {
		t.Fatalf(`Postcode mapped improperly`)
	}

	if employee.Email != "test@email.com" {
		t.Fatalf(`Email mapped improperly`)
	}

	if employee.StartDate != startDate {
		t.Fatalf(`StartDate mapped improperly`)
	}
}

func Test_ExistingEmployee_FieldsMappedProperly(t *testing.T) {

	dateOfBirth, _ := time.Parse(time.DateOnly, "1990-12-11")
	startDate, _ := time.Parse(time.DateOnly, "2025-10-24")
	id := uuid.New()
	employee := employees.ExistingEmployee(
		id,
		"TEST_PayrollNumber",
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

	if employee.Id != id {
		t.Fatalf(`Id mapped improperly`)
	}

	if employee.PayrollNumber != "TEST_PayrollNumber" {
		t.Fatalf(`PayrollNumber mapped improperly`)
	}

	if employee.Forenames != "TEST_Forenames" {
		t.Fatalf(`Forenames mapped improperly`)
	}

	if employee.Surname != "TEST_Surname" {
		t.Fatalf(`Surname mapped improperly`)
	}

	if employee.DateOfBirth != dateOfBirth {
		t.Fatalf(`DateOfBirth mapped improperly`)
	}

	if employee.TelephoneNumber != "998901234567" {
		t.Fatalf(`TelephoneNumber mapped improperly`)
	}

	if employee.MobileNumber != "998129834567" {
		t.Fatalf(`MobileNumber mapped improperly`)
	}

	if employee.AddressLine1 != "TEST_AddressLine1" {
		t.Fatalf(`AddressLine1 mapped improperly`)
	}

	if employee.AddressLine2 != "TEST_AddressLine2" {
		t.Fatalf(`AddressLine2 mapped improperly`)
	}

	if employee.Postcode != "TEST_Postcode" {
		t.Fatalf(`Postcode mapped improperly`)
	}

	if employee.Email != "test@email.com" {
		t.Fatalf(`Email mapped improperly`)
	}

	if employee.StartDate != startDate {
		t.Fatalf(`StartDate mapped improperly`)
	}
}
