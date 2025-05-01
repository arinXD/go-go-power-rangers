package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var books []Book

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load .env error")
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	books = append(books, Book{Id: 1, Title: "Go Go Power Ranger", Author: Author{Name: "Arin Lnwza007", Age: 14}})

	app.Use(logger)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Listen(os.Getenv("PORT"))
}

