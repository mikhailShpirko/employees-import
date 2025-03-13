package tests

import (
	"employees-import/features/employees"
	parser "employees-import/parsers"
	"os"
	"slices"
	"testing"
	"time"
)

func Test_ParseEmployeeDataFromCsv(t *testing.T) {
	file, fileErr := os.Open("test_employee_import.csv")

	if fileErr != nil {
		t.Fatalf(`Error opening test file %v`, fileErr)
	}

	result, parseErr := parser.ParseEmployeeDataFromCsv(file)

	closeErr := file.Close()

	if closeErr != nil {
		t.Fatalf(`Error closing test file %v`, closeErr)
	}

	if parseErr != nil {
		t.Fatalf(`Error parsing test file %v`, fileErr)
	}

	switch parseResult := result.(type) {
	case parser.EmployeesSuccessfullyParsed:

		if !slices.Equal(parseResult.Result, expectedEmployees()) {
			t.Fatalf(`Parsed employees don't matched expected %v`, parseResult.Result)
		}

	case parser.ErrorParsingEmployees:
		t.Fatalf("Unexpected parsing errors %v", parseResult.ErrorInLines)
	default:
		t.Fatalf("Unsupported result")
	}
}

func expectedEmployees() []employees.EmployeeData {

	employee1 := employees.EmployeeData{}
	employee1.PayrollNumber = "COOP08"
	employee1.Forenames = "John"
	employee1.Surname = "William"
	employee1.DateOfBirth, _ = time.Parse("2/1/2006", "26/01/1955")
	employee1.TelephoneNumber = "12345678"
	employee1.MobileNumber = "987654231"
	employee1.AddressLine1 = "12 Foreman road"
	employee1.AddressLine2 = "London"
	employee1.Postcode = "GU12 6JW"
	employee1.Email = "nomadic20@hotmail.co.uk"
	employee1.StartDate, _ = time.Parse("2/1/2006", "18/04/2013")

	employee2 := employees.EmployeeData{}
	employee2.PayrollNumber = "JACK13"
	employee2.Forenames = "Jerry"
	employee2.Surname = "Jackson"
	employee2.DateOfBirth, _ = time.Parse("2/1/2006", "11/5/1974")
	employee2.TelephoneNumber = "2050508"
	employee2.MobileNumber = "6987457"
	employee2.AddressLine1 = "115 Spinney Road"
	employee2.AddressLine2 = "Luton"
	employee2.Postcode = "LU33DF"
	employee2.Email = "gerry.jackson@bt.com"
	employee2.StartDate, _ = time.Parse("2/1/2006", "18/04/2013")

	return []employees.EmployeeData{employee1, employee2}
}
