package endpoints_employees

import (
	endpoints_employees_create "employees-import/endpoints/employees/create"
	endpoints_employees_delete "employees-import/endpoints/employees/delete"
	endpoints_employees_getAll "employees-import/endpoints/employees/getAll"
	endpoints_employees_getById "employees-import/endpoints/employees/getById"
	endpoints_employees_import "employees-import/endpoints/employees/import"
	endpoints_employees_update "employees-import/endpoints/employees/update"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeesEndpoints(app *fiber.App) {

	employeesEndpoints := app.Group("/employees")

	employeesEndpoints.Post("/import/csv", endpoints_employees_import.HandleRequest)

	employeesEndpoints.Get("/", endpoints_employees_getAll.HandleRequest)
	employeesEndpoints.Get("/:id", endpoints_employees_getById.HandleRequest)

	employeesEndpoints.Post("/", endpoints_employees_create.HandleRequest)
	employeesEndpoints.Put("/:id", endpoints_employees_update.HandleRequest)
	employeesEndpoints.Delete("/:id", endpoints_employees_delete.HandleRequest)
}
