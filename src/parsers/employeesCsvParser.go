package parsers

import (
	employees "employees-import/features/employees"
	"io"
	"slices"
	"strings"
	"time"

	"encoding/csv"
)

func ParseEmployeeDataFromCsv(fileReader io.Reader) (IParseEmployeesResult, error) {
	lines, readErr := csv.NewReader(fileReader).ReadAll()

	if readErr != nil {
		return nil, readErr
	}

	result := []employees.EmployeeData{}
	errors := map[int]EMPLOYEE_PARSING_ERROR{}

	for i, line := range lines {

		if len(line) != 11 {
			errors[i+1] = EMPLOYEE_PARSING_ERROR_INVALID_LINE_FORMAT
			continue
		}

		if i == 0 && isHeaderLine(line) {
			continue
		}

		dateOfBirth, dateOfBirthPareError := time.Parse("2/1/2006", line[3])

		if dateOfBirthPareError != nil {
			errors[i+1] = EMPLOYEE_PARSING_ERROR_INVALID_DATE_OF_BIRTH
			continue
		}

		startDate, startDatePareError := time.Parse("2/1/2006", line[10])

		if startDatePareError != nil {
			errors[i+1] = EMPLOYEE_PARSING_ERROR_INVALID_START_DATE
			continue
		}

		employee := employees.EmployeeData{}

		employee.PayrollNumber = strings.TrimSpace(line[0])
		employee.Forenames = strings.TrimSpace(line[1])
		employee.Surname = strings.TrimSpace(line[2])
		employee.DateOfBirth = dateOfBirth
		employee.TelephoneNumber = strings.TrimSpace(line[4])
		employee.MobileNumber = strings.TrimSpace(line[5])
		employee.AddressLine1 = strings.TrimSpace(line[6])
		employee.AddressLine2 = strings.TrimSpace(line[7])
		employee.Postcode = strings.TrimSpace(line[8])
		employee.Email = strings.TrimSpace(line[9])
		employee.StartDate = startDate

		result = append(result, employee)
	}

	if len(errors) > 0 {
		return ErrorParsingEmployees{ErrorInLines: errors}, nil
	}

	return EmployeesSuccessfullyParsed{Result: result}, nil
}

type IParseEmployeesResult interface{}

type BaseParseEmployeesResult struct{}

type EmployeesSuccessfullyParsed struct {
	BaseParseEmployeesResult
	Result []employees.EmployeeData
}

type ErrorParsingEmployees struct {
	BaseParseEmployeesResult
	ErrorInLines map[int]EMPLOYEE_PARSING_ERROR
}

type EMPLOYEE_PARSING_ERROR int

const (
	EMPLOYEE_PARSING_ERROR_INVALID_LINE_FORMAT EMPLOYEE_PARSING_ERROR = iota
	EMPLOYEE_PARSING_ERROR_INVALID_DATE_OF_BIRTH
	EMPLOYEE_PARSING_ERROR_INVALID_START_DATE
)

func (validationError EMPLOYEE_PARSING_ERROR) String() string {
	return [...]string{"INVALID_LINE_FORMAT",
		"INVALID_DATE_OF_BIRTH",
		"INVALID_START_DATE"}[validationError]
}

var headerLine = []string{
	"Personnel_Records.Payroll_Number",
	"Personnel_Records.Forenames",
	"Personnel_Records.Surname",
	"Personnel_Records.Date_of_Birth",
	"Personnel_Records.Telephone",
	"Personnel_Records.Mobile",
	"Personnel_Records.Address",
	"Personnel_Records.Address_2",
	"Personnel_Records.Postcode",
	"Personnel_Records.EMail_Home",
	"Personnel_Records.Start_Date"}

func isHeaderLine(line []string) bool {
	return slices.Equal(line, headerLine)
}
