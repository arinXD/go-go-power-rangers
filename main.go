package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var books []Book

func main() {
	books = append(books, Book{Id: 1, Title: "Go Go Power Ranger", Author: Author{Name: "Arin Lnwza007", Age: 14}})
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load .env error")
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger)
	app.Post("/login", login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	app.Use(IsAdmin)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Listen(os.Getenv("PORT"))
}
