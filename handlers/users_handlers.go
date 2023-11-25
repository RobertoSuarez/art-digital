package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// metodos de esta estructura
func (userHandler *UserHandler) PostRegisterUser(c *fiber.Ctx) error {
	var apiUser types.APIUser

	if err := c.BodyParser(&apiUser); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Error en los datos",
			})
	}

	err := data.RegisterUser(&apiUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "No se pudo guardar los datos",
			})
	}

	return c.Status(fiber.StatusOK).JSON(apiUser)
}

func (userHandler *UserHandler) PostLogin(c *fiber.Ctx) error {
	var login types.Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Error en los datos",
			})
	}

	user, ok, err := data.Login(login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := userHandler.createToken(user)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user":  user,
	})

}

type CustonClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (userHandler *UserHandler) createToken(user types.APIUser) (string, error) {

	claims := CustonClaims{
		user.ID,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    user.Type,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := viper.GetString("APP_SECRET_TOKEN")

	return token.SignedString([]byte(secret))
}

// VerifiToken es un midleware para validar el token.
func (userHandler *UserHandler) VerifiToken(c *fiber.Ctx) error {

	// Recuperamos el token.
	auth := c.Get("Authorization")
	if auth == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// separamos el token.
	arrayToken := strings.Split(auth, " ")
	tokenString := arrayToken[1]

	// verificamos el token y recuperamos los claims.
	token, err := jwt.ParseWithClaims(tokenString, &CustonClaims{}, func(t *jwt.Token) (interface{}, error) {
		secret := viper.GetString("APP_SECRET_TOKEN")
		return []byte(secret), nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(*CustonClaims)
	if ok {
		fmt.Println(claims)
		c.Locals("claims", claims)
	}

	return c.Next()
}
