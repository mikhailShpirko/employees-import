package endpoints_employees_import

import (
	configuration "employees-import/configuration"
	employees_import "employees-import/features/employees/import"
	parsers "employees-import/parsers"
	persistence "employees-import/persistence"

	"github.com/gofiber/fiber/v2"
)

func HandleRequest(c *fiber.Ctx) error {

	fileHeader, fileHeaderErr := c.FormFile("import-file")

	if fileHeaderErr != nil {
		return fileHeaderErr
	}

	if fileHeader.Header.Get("Content-Type") != "text/csv" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Only .csv files allowed")
	}

	file, fileErr := fileHeader.Open()

	if fileErr != nil {
		return fileErr
	}

	result, _ := parsers.ParseEmployeeDataFromCsv(file)

	fileCloseErr := file.Close()

	if fileCloseErr != nil {
		return fileCloseErr
	}

	switch parseResult := result.(type) {
	case parsers.EmployeesSuccessfullyParsed:
		repository, unitOfWork, persistenceError := persistence.CreateEmployeeRepository(configuration.GetDatabaseConnectionString(), c.Context())

		if persistenceError != nil {
			return persistenceError
		}

		defer repository.CloseConnection()
		defer unitOfWork.Rollback()

		importResult, importError := employees_import.Handle(parseResult.Result, repository, unitOfWork)

		if importError != nil {
			return importError
		}

		switch result := importResult.(type) {
		case employees_import.SuccessfullyImported:
			response := []EmployeeImportResult{}
			for _, employeeImportResult := range result.Result {
				response = append(response, EmployeeImportResult{PayrollNumber: employeeImportResult.PayrollNumber, StatusCode: employeeImportResult.Status.String(), Id: employeeImportResult.Id})
			}
			return c.Status(fiber.StatusOK).JSON(response)
		case employees_import.ValidationErrors:
			response := []ErrorImportingEmployee{}
			for _, validationErrors := range result.Errors {
				errorCodes := []string{}
				for _, validationError := range validationErrors.ValidationErrors {
					errorCodes = append(errorCodes, validationError.String())
				}
				response = append(response, ErrorImportingEmployee{PayrollNumber: validationErrors.PayrollNumber, ErrorCodes: errorCodes})
			}
			return c.Status(fiber.StatusBadRequest).JSON(response)
		default:
			return fiber.NewError(fiber.StatusInternalServerError, "Unsupported import result")
		}
	case parsers.ErrorParsingEmployees:
		response := []ErrorParsingLine{}
		for lineNumber, errorCode := range parseResult.ErrorInLines {
			response = append(response, ErrorParsingLine{LineNumber: lineNumber, ErrorCode: errorCode.String()})
		}
		return c.Status(fiber.StatusNotAcceptable).JSON(response)
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "Unsupported parsing result")
	}
}
