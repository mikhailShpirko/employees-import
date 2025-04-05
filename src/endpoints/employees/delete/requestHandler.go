package endpoints_employees_delete

import (
	configuration "employees-import/configuration"
	employees_delete "employees-import/features/employees/delete"
	persistence "employees-import/persistence"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleRequest(c *fiber.Ctx) error {

	id, parseError := uuid.Parse(c.Params("id"))

	if parseError != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id must be a valid UUID")
	}

	repository, unitOfWork, persistenceError := persistence.CreateEmployeeRepository(configuration.GetDatabaseConnectionString(), c.Context())

	if persistenceError != nil {
		return persistenceError
	}

	defer repository.CloseConnection()
	defer unitOfWork.Rollback()

	result, err := employees_delete.Handle(id, repository, unitOfWork)

	if err != nil {
		return err
	}

	return employees_delete.Match(result,
		func(_ employees_delete.Deleted) error {
			return c.Status(fiber.StatusNoContent).JSON("")
		},
		func(employeeNotExists employees_delete.EmployeeNotExists) error {
			return c.Status(fiber.StatusNotFound).JSON("Employee not found")
		})
}
