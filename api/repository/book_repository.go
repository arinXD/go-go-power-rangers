package repository

import (
	"github.com/arinxd/gogo/api/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (r *BookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	result := r.DB.Preload("Author").Find(&books)
	return books, result.Error
}

func (r *BookRepository) FindByID(id int) (*models.Book, error) {
	var book models.Book
	result := r.DB.Preload("Author").First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *BookRepository) Create(book *models.Book) error {
	return r.DB.Create(book).Error
}
func (r *BookRepository) Update(book *models.Book, data map[string]interface{}) error {
	return r.DB.Model(book).Updates(data).Error
}

func (r *BookRepository) Delete(book *models.Book) error {
	return r.DB.Delete(book).Error
}
func (r *BookRepository) Exists(id int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Book{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}