package main

import (
	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.Init()

	db.DB.AutoMigrate(
		&data.User{},
		&data.Country{},
		&data.Art{},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola, mundo")
	})

	app.Listen(":3000")
}
