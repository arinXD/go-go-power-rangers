package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func httpLogger(c *fiber.Ctx) error {
	url := c.OriginalURL()
	method := c.Method()
	log.Printf("[%s] %s", method, url)
	return c.Next()
}