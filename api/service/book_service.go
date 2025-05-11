package service

import (
	"errors"
	"github.com/arinxd/gogo/api/models"
	"github.com/arinxd/gogo/api/repository"
	"gorm.io/gorm"
)

type BookService struct {
	BookRepo *repository.BookRepository
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		BookRepo: repository.NewBookRepository(db),
	}
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.BookRepo.FindAll()
}
func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.BookRepo.FindByID(id)
}

func (s *BookService) CreateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}
	if book.ISBN == "" {
		return errors.New("book ISBN cannot be empty")
	}
	
	return s.BookRepo.Create(book)
}

func (s *BookService) UpdateBook(id int, bookData *models.Book) (*models.Book, error) {
	book, err := s.BookRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	
	updates := map[string]interface{}{
		"title":        bookData.Title,
		"isbn":         bookData.ISBN,
		"description":  bookData.Description,
		"published_at": bookData.PublishedAt,
		"author_id":    bookData.AuthorID,
	}
	
	if bookData.Title == "" {
		return nil, errors.New("book title cannot be empty")
	}
	
	if err := s.BookRepo.Update(book, updates); err != nil {
		return nil, err
	}
	
	return s.BookRepo.FindByID(id)
}

func (s *BookService) DeleteBook(id int) error {
	exists, err := s.BookRepo.Exists(id)
	if err != nil {
		return err
	}
	if !exists {
		return gorm.ErrRecordNotFound
	}
	
	book, err := s.BookRepo.FindByID(id)
	if err != nil {
		return err
	}
	
	return s.BookRepo.Delete(book)
}