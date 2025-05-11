package repository

import (
	"github.com/arinxd/gogo/api/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		DB: db,
	}
}

func (r *AuthorRepository) CreateAuthor(author *models.Author) error {
	return r.DB.Create(author).Error
}