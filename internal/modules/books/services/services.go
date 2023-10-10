package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/repositories"

	"github.com/google/uuid"
)

type BookServicesImpl interface {
	Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookResponse, error)
	Show(id uuid.UUID) (*dto.BookResponse, error)
	Index(limit, page int64) (*dto.BookIndexResponse, error)
	Update(id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error)
	Delete(id *uuid.UUID) error
}

type BookService struct {
	BookRepository repositories.BookRepositoryImpl
}

func NewBookService(
	bookRepository repositories.BookRepositoryImpl,
) BookServicesImpl {

	return &BookService{bookRepository}
}
