package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var user = User{
	Email:    "admin@admin.com",
	Password: "admin",
}

func isAuth(c *fiber.Ctx) error {
	loginUser := new(User)
	if err := c.BodyParser(loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if user.Email == loginUser.Email && user.Password == loginUser.Password {
		return c.Next()
	}
	return fiber.ErrUnauthorized
}
