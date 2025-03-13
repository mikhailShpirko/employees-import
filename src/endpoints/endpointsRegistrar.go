package endpoints

import (
	employees "employees-import/endpoints/employees"

	"github.com/gofiber/fiber/v2"
)

func RegisterEndpoints(app *fiber.App) {
	employees.RegisterEmployeesEndpoints(app)
}
