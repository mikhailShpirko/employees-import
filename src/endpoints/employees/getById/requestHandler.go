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

	result, getByIdError := employees_getById.Handle(id, repository)

	if getByIdError != nil {
		return getByIdError
	}

	return employees_getById.Match(result,
		func(exists employees_getById.EmployeeExists) error {
			response := EmployeeData{}
			response.PayrollNumber = exists.Employee.PayrollNumber
			response.Forenames = exists.Employee.Forenames
			response.Surname = exists.Employee.Surname
			response.DateOfBirth = custom_types.DateOnly{Time: exists.Employee.DateOfBirth}
			response.TelephoneNumber = exists.Employee.TelephoneNumber
			response.MobileNumber = exists.Employee.MobileNumber
			response.AddressLine1 = exists.Employee.AddressLine1
			response.AddressLine2 = exists.Employee.AddressLine2
			response.Postcode = exists.Employee.Postcode
			response.Email = exists.Employee.Email
			response.StartDate = custom_types.DateOnly{Time: exists.Employee.StartDate}

			return c.Status(fiber.StatusOK).JSON(response)
		},
		func(notExists employees_getById.EmployeeNotExists) error {
			return c.Status(fiber.StatusNotFound).JSON("Employee not found")
		})
}
