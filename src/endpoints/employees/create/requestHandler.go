package endpoints_employees_create

import (
	configuration "employees-import/configuration"
	employees "employees-import/features/employees"
	employees_create "employees-import/features/employees/create"
	persistence "employees-import/persistence"

	"github.com/gofiber/fiber/v2"
)

func HandleRequest(c *fiber.Ctx) error {

	dto := CreateEmployee{}

	err := c.BodyParser(&dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorCreatingEmployee{ErrorCodes: []string{"INVALID_DATA"}})
	}

	repository, unitOfWork, persistenceError := persistence.CreateEmployeeRepository(configuration.GetDatabaseConnectionString(), c.Context())

	if persistenceError != nil {
		return persistenceError
	}

	defer repository.CloseConnection()
	defer unitOfWork.Rollback()

	result, err := employees_create.Handle(employees.CreateEmployeeData(dto.PayrollNumber,
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

	switch createResult := result.(type) {
	case employees_create.Created:
		return c.Status(fiber.StatusCreated).JSON(EmployeeCreated{Id: createResult.Id})
	case employees_create.PayrollNumberAlreadyExists:
		return fiber.NewError(fiber.StatusConflict, "Payroll number already exists")
	case employees_create.ValidationErrors:
		errorCodes := []string{}
		for _, errorCode := range createResult.Errors {
			errorCodes = append(errorCodes, errorCode.String())
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorCreatingEmployee{ErrorCodes: errorCodes})
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "Unsupported create result")
	}
}
