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

	defer unitOfWork.Rollback()

	result, err := employees_delete.Handle(id, repository, unitOfWork)

	if err != nil {
		return err
	}

	switch result.(type) {
	case employees_delete.Deleted:
		return c.Status(fiber.StatusNoContent).JSON("")
	case employees_delete.EmployeeNotExists:
		return fiber.NewError(fiber.StatusNotFound, "Employee not found")
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "Unsupported delete result")
	}
}
