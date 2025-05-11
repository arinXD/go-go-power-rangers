package handler

import (
	"log"

	"github.com/arinxd/gogo/api/models"
	"github.com/arinxd/gogo/api/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthorHandler struct {
	AuthorService *service.AuthorService
}

func NewAuthorHandler(db *gorm.DB) *AuthorHandler {
	return &AuthorHandler{
		AuthorService: service.NewAuthorService(db),
	}
}

func (h *AuthorHandler) CreateAuthor(c *fiber.Ctx) error {
	author := new(models.Author)
	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	if err := h.AuthorService.CreateAuthor(author); err != nil {
		log.Println(author)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create author",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}
