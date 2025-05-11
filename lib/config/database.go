package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

var (
	host       string
	port       string
	dbUser     string
	dbPassword string
	dbName     string
	dbLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
)

func InitDBConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	dbUser = os.Getenv("POSTGRES_USER")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("POSTGRES_DB")
}

func GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPassword, dbName)
}

func GetDbLogger() logger.Interface {
	return dbLogger
}