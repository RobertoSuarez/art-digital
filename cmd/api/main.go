package main

import (
	"fmt"

	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/db"
	"github.com/RobertoSuarez/art-digital/handlers"
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

	api := app.Group("/api")

	users := api.Group("/users")
	userHandlers := handlers.NewUserHandler()

	users.Post("/", userHandlers.PostRegisterUser)
	users.Post("/login", userHandlers.PostLogin)

	api.Get("/protegida", userHandlers.VerifiToken, func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*handlers.CustonClaims)
		if ok {
			fmt.Println(claims)
		}
		return c.SendString("Si tenemos el token")
	})

	app.Listen(":4000")
}
