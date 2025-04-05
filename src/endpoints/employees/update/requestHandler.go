package endpoints_employees_update

import (
	configuration "employees-import/configuration"
	employees "employees-import/features/employees"
	employees_update "employees-import/features/employees/update"
	persistence "employees-import/persistence"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleRequest(c *fiber.Ctx) error {

	id, parseError := uuid.Parse(c.Params("id"))

	if parseError != nil {

		return c.Status(fiber.StatusBadRequest).JSON(ErrorUpdatingEmployee{ErrorCodes: []string{"INVALID_ID"}})
	}

	dto := UpdateEmployee{}

	err := c.BodyParser(&dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorUpdatingEmployee{ErrorCodes: []string{"INVALID_DATA"}})
	}

	repository, unitOfWork, persistenceError := persistence.CreateEmployeeRepository(configuration.GetDatabaseConnectionString(), c.Context())

	if persistenceError != nil {
		return persistenceError
	}

	defer repository.CloseConnection()
	defer unitOfWork.Rollback()

	result, err := employees_update.Handle(employees.CreateEmployee(id,
		dto.PayrollNumber,
		dto.Forenames,
		dto.Surname,
		dto.DateOfBirth.Time,
		dto.TelephoneNumber,
		dto.MobileNumber,
		dto.AddressLine1,
		dto.AddressLine2,
		dto.Postcode,
		dto.Email,
		dto.StartDate.Time),
		repository,
		unitOfWork)

	if err != nil {
		return err
	}

	return employees_update.Match(result,
		func(updated employees_update.Updated) error {
			return c.Status(fiber.StatusNoContent).JSON("")
		},
		func(payrollExists employees_update.PayrollNumberAlreadyExists) error {
			return c.Status(fiber.StatusConflict).JSON("Payroll number already exists")
		},
		func(validationErrors employees_update.ValidationErrors) error {
			errorCodes := []string{}
			for _, errorCode := range validationErrors.Errors {
				errorCodes = append(errorCodes, errorCode.String())
			}
			return c.Status(fiber.StatusBadRequest).JSON(ErrorUpdatingEmployee{ErrorCodes: errorCodes})
		},
		func(notExists employees_update.EmployeeNotExists) error {
			return c.Status(fiber.StatusNotFound).JSON("Employee not found")
		})
}
