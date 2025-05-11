package service

import (
	"errors"
	"github.com/arinxd/gogo/api/models"
	"github.com/arinxd/gogo/api/repository"
	"gorm.io/gorm"
)

type AuthorService struct {
	AuthorRepo *repository.AuthorRepository
}

func NewAuthorService(db *gorm.DB) *AuthorService {
	return &AuthorService{
		AuthorRepo: repository.NewAuthorRepository(db),
	}
}

func (s *AuthorService) CreateAuthor(author *models.Author) error {
	if author.Name == "" {
		return errors.New("author name cannot be empty")
	}
	if author.Biography == "" {
		return errors.New("author biography cannot be empty")
	}
	
	return s.AuthorRepo.CreateAuthor(author)
}	