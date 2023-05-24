package book

import "phsy_rsv_go/domain"

type Service interface {
	FindAll() ([]domain.Book, error)
	FindByID(ID int) (domain.Book, error)
	Create(bookrequest domain.BookRequest) (domain.Book, error)
	Update(ID int, bookrequest domain.BookRequest) (domain.Book, error)
	Delete(ID int) (domain.Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.Book, error) {

	books, err := s.repository.FindAll()

	return books, err
}

func (s *service) FindByID(ID int) (domain.Book, error) {
	book, err := s.repository.FindByID(ID)

	return book, err
}

func (s *service) Create(bookrequest domain.BookRequest) (domain.Book, error) {

	price, _ := bookrequest.Price.Int64()

	book, err := s.repository.Create(domain.Book{
		Title: bookrequest.Title,
		Price: int(price),
	})

	return book, err
}

func (s *service) Update(ID int, bookrequest domain.BookRequest) (domain.Book, error) {

	book, _ := s.repository.FindByID(ID)

	price, _ := bookrequest.Price.Int64()

	book.Price = int(price)
	book.Title = bookrequest.Title

	newbook, err := s.repository.Update(book)

	return newbook, err
}

func (s *service) Delete(ID int) (domain.Book, error) {
	book, _ := s.repository.FindByID(ID)
	deletedBook, err := s.repository.Delete(book)

	return deletedBook, err
}
