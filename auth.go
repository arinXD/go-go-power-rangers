package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func hasRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] == role || claims["role"] == "admin" {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "No permission",
			"status": fiber.StatusForbidden,
		})
	}
}

var IsAdmin = hasRole("admin")
var IsMember = hasRole("member")
