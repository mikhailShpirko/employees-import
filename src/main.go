package main

import (
	endpoints "employees-import/endpoints"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	fiberLog "github.com/gofiber/fiber/v2/log"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	//.env file is optional. so ignore error if file not exists
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: handleError,
	})

	app.Use(fiberRecover.New())

	endpoints.RegisterEndpoints(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func handleError(c *fiber.Ctx, err error) error {
	m := ErrorResponse{
		Status:  fiber.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	if e, ok := err.(*fiber.Error); ok {
		m.Status = e.Code
		m.Message = e.Message
	} else {
		fiberLog.Error(`Unexpected error`, err)
	}

	return c.Status(m.Status).JSON(m)
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
