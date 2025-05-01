package main

import (
	"os"
	"path/filepath"
	"time"
	"github.com/gofiber/fiber/v2"
)

func uploadFile(c *fiber.Ctx) error {
	if err := os.MkdirAll("./upload", 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	timestamp := time.Now().Format("20060102150405")
	fileExt := filepath.Ext(file.Filename)
	newFilename := "./upload/" + timestamp + fileExt
	err = c.SaveFile(file, newFilename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("Upload file completed.")
}
