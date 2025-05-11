package main

import (
	"os"
	"github.com/arinxd/gogo/api/handler"
	"github.com/arinxd/gogo/lib/config"
	"github.com/arinxd/gogo/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/jwt/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	config.InitDBConfig()
	var err error
	DB := &gorm.DB{}
	dsn := config.GetDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: config.GetDbLogger(),
	})
	if err != nil {
		panic("failed to connect to database")
	}
	DB.AutoMigrate(&models.Author{}, &models.Book{})

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(httpLogger)

	api := app.Group("/api/v1")
	api.Post("/login", login)
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	bookHandler := handler.NewBookHandler(DB)
	authorHandler := handler.NewAuthorHandler(DB)

	api.Use(IsAdmin)
	bookGroup := api.Group("/books")
	bookGroup.Get("", bookHandler.GetBooks)
	bookGroup.Get("/:id", bookHandler.GetBookById)
	bookGroup.Post("/", bookHandler.CreateBook)
	bookGroup.Put("/:id", bookHandler.UpdateBook)
	bookGroup.Delete("/:id", bookHandler.DeleteBook)

	authorGroup := api.Group("/authors")
	authorGroup.Post("/", authorHandler.CreateAuthor)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}