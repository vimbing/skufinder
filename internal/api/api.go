package api

import "github.com/gofiber/fiber/v2"

func Api() {
	app := fiber.New()

	app.Post("/find", Find)

	app.Listen(":3000")
}
