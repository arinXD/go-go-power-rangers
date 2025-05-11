package models

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	gorm.Model
	Name        string    `json:"name"`
	Biography   string    `json:"biography"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Books       []Book    `json:"books" gorm:"foreignKey:AuthorID"`
}

type Book struct {
	gorm.Model
	Title       string    `json:"title"`
	ISBN        string    `json:"isbn" gorm:"uniqueIndex"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	AuthorID    uint      `json:"author_id"`
	Author      Author    `json:"author" gorm:"foreignKey:AuthorID"`
}