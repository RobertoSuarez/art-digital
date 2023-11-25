package handlers

import (
	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/types"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type UserController struct {
	viper *viper.Viper
}

func NewUserController(viper *viper.Viper) *UserController {
	return &UserController{
		viper: viper,
	}
}

func (userc *UserController) HandlerRegisterUser(c *fiber.Ctx) error {

	var apiUser types.UserAPI
	if err := c.BodyParser(&apiUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := data.RegisterUser(&apiUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(apiUser)

}
