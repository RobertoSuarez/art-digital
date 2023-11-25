package main

import (
	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/db"
	"github.com/RobertoSuarez/art-digital/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	viperApp := viper.New()

	viperApp.AutomaticEnv()

	app := fiber.New()

	db.Init()

	db.DB.AutoMigrate(
		&data.Country{},
		&data.Art{},
		&data.User{},
	)

	db.DB.Where(&data.Country{Name: "Ecuador"}).
		FirstOrCreate(&data.Country{
			Name:          "Ecuador",
			Code:          "EC",
			ContinentName: "America",
		})

	api := app.Group("/api")

	users := api.Group("/users")

	userController := handlers.NewUserController(viperApp)

	// registrar un usuario
	users.Post("/", userController.HandlerRegisterUser)

	users.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("login correcto")
	})

	app.Listen(":4000")
}
