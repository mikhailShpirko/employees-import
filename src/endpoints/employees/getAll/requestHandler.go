package endpoints_employees_getAll

import (
	configuration "employees-import/configuration"
	custom_types "employees-import/customTypes"
	employees_getAll "employees-import/features/employees/getAll"
	persistence "employees-import/persistence"

	"github.com/gofiber/fiber/v2"
)

func HandleRequest(c *fiber.Ctx) error {

	repository, _, persistenceError := persistence.CreateEmployeeRepository(configuration.GetDatabaseConnectionString(), c.Context())

	if persistenceError != nil {
		return persistenceError
	}

	defer repository.CloseConnection()

	employees, getAllError := employees_getAll.Handle(repository)

	if getAllError != nil {
		return getAllError
	}

	response := []Employee{}

	for _, employee := range employees {
		dto := Employee{}
		dto.Id = employee.Id
		dto.PayrollNumber = employee.PayrollNumber
		dto.Forenames = employee.Forenames
		dto.Surname = employee.Surname
		dto.DateOfBirth = custom_types.DateOnly{Time: employee.DateOfBirth}
		dto.TelephoneNumber = employee.TelephoneNumber
		dto.MobileNumber = employee.MobileNumber
		dto.AddressLine1 = employee.AddressLine1
		dto.AddressLine2 = employee.AddressLine2
		dto.Postcode = employee.Postcode
		dto.Email = employee.Email
		dto.StartDate = custom_types.DateOnly{Time: employee.StartDate}

		response = append(response, dto)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
