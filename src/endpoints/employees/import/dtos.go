package endpoints_employees_import

import "github.com/google/uuid"

type EmployeeImportResult struct {
	PayrollNumber string    `json:"payrolNumber"`
	StatusCode    string    `json:"statusCode"`
	Id            uuid.UUID `json:"id"`
}

type ErrorImportingEmployee struct {
	PayrollNumber string   `json:"payrolNumber"`
	ErrorCodes    []string `json:"errorCodes"`
}

type ErrorParsingLine struct {
	LineNumber int    `json:"lineNumber"`
	ErrorCode  string `json:"errorCode"`
}
