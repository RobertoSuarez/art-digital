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
		&data.Country{},
		&data.Art{},
		&data.User{},
	)

	app.Listen(":3000")
}
