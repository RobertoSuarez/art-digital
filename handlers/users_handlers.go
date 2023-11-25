package handlers

import (
	"strings"
	"time"

	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func (userc *UserController) HandlerLogin(c *fiber.Ctx) error {

	var login types.Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error en los datos",
		})
	}

	userAPI, err := data.Login(&login)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	secret := []byte(userc.viper.GetString("JWT_SECRET"))

	claims := types.UserClaims{
		ID:     userAPI.ID,
		Name:   userAPI.Name,
		Email:  userAPI.Email,
		Status: userAPI.Status,
		Type:   userAPI.Type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			Issuer:    userAPI.Type,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
		"user":  userAPI,
	})

}

func (userc *UserController) JWT(c *fiber.Ctx) error {
	auth := c.Get("Authorization")

	authSlice := strings.Split(auth, " ")
	if len(authSlice) < 1 {
		return c.Status(fiber.StatusUnauthorized).SendString("No esta autorizado")
	}

	tokenString := authSlice[1]

	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		secret := []byte(userc.viper.GetString("JWT_SECRET"))
		return secret, nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(*types.UserClaims)

	if ok {
		c.Locals("claims", claims)
		return c.Next()
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

}
