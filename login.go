package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Email string
	Role  string
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var user = UserLogin{
	Email:    "admin@admin.com",
	Password: "admin",
}

func login(c *fiber.Ctx) error {
	userLogin := new(UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if userLogin.Email != user.Email && userLogin.Password != user.Password {
		return fiber.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userLogin.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"email": user.Email,
		"token": t,
	})
}
