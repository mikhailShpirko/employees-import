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

	defer unitOfWork.Rollback()

	result, err := employees_update.Handle(employees.ExistingEmployee(id,
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

	switch updateResult := result.(type) {
	case employees_update.Updated:
		return c.Status(fiber.StatusNoContent).JSON("")
	case employees_update.PayrollNumberAlreadyExists:
		return fiber.NewError(fiber.StatusConflict, "Payroll number already exists")
	case employees_update.ValidationErrors:
		errorCodes := []string{}
		for _, errorCode := range updateResult.Errors {
			errorCodes = append(errorCodes, errorCode.String())
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorUpdatingEmployee{ErrorCodes: errorCodes})
	case employees_update.EmployeeNotExists:
		return fiber.NewError(fiber.StatusNotFound, "Employee not found")
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "Unsupported update result")
	}
}
