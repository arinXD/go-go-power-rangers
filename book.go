package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if id == book.Id {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("This book ID not found.")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book) ; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdated := new(Book)
	if err := c.BodyParser(bookUpdated) ; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	
	for i, book := range books {
		if id == book.Id {
			books[i].Title = bookUpdated.Title
			books[i].Author = bookUpdated.Author
			return c.JSON(books[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("This book ID not found.")
}

func deleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if id == book.Id {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("This book ID not found.")
}
