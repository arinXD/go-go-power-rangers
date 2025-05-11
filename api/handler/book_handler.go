package handler

import (
	"log"
	"strconv"

	"github.com/arinxd/gogo/api/models"
	"github.com/arinxd/gogo/api/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookHandler struct {
	BookService *service.BookService
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		BookService: service.NewBookService(db),
	}
}

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.BookService.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch books",
		})
	}
	return c.JSON(books)
}

func (h *BookHandler) GetBookById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	book, err := h.BookService.GetBookByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch book",
		})
	}

	return c.JSON(book)
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	
	if err := h.BookService.CreateBook(book); err != nil {
		log.Println(book)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create book",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	updatedBook := new(models.Book)
	if err := c.BodyParser(updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	book, err := h.BookService.UpdateBook(id, updatedBook)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update book",
		})
	}

	return c.JSON(book)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	err = h.BookService.DeleteBook(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete book",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}