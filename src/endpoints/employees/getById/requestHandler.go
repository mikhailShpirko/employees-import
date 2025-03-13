package endpoints_employees_getById

import (
	configuration "employees-import/configuration"
	custom_types "employees-import/customTypes"
	employees_getById "employees-import/features/employees/getById"
	persistence "employees-import/persistence"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleRequest(c *fiber.Ctx) error {

	id, parseError := uuid.Parse(c.Params("id"))

	if parseError != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id must be a valid UUID")
	}

	repository, _, persistenceError := persistence.CreateEmployeeRepository(configuration.GetDatabaseConnectionString(), c.Context())

	if persistenceError != nil {
		return persistenceError
	}

	defer repository.CloseConnection()

	employeeResult, getByIdError := employees_getById.Handle(id, repository)

	if getByIdError != nil {
		return getByIdError
	}

	switch result := employeeResult.(type) {
	case employees_getById.EmployeeExists:
		response := EmployeeData{}
		response.PayrollNumber = result.Employee.PayrollNumber
		response.Forenames = result.Employee.Forenames
		response.Surname = result.Employee.Surname
		response.DateOfBirth = custom_types.DateOnly{Time: result.Employee.DateOfBirth}
		response.TelephoneNumber = result.Employee.TelephoneNumber
		response.MobileNumber = result.Employee.MobileNumber
		response.AddressLine1 = result.Employee.AddressLine1
		response.AddressLine2 = result.Employee.AddressLine2
		response.Postcode = result.Employee.Postcode
		response.Email = result.Employee.Email
		response.StartDate = custom_types.DateOnly{Time: result.Employee.StartDate}

		return c.Status(fiber.StatusOK).JSON(response)
	case employees_getById.EmployeeNotExists:
		return fiber.NewError(fiber.StatusNotFound, "Employee not found")
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "Unsupported get by id result")
	}
}
