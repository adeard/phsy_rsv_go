package book

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.Book, error)
	FindByID(ID int) (domain.Book, error)
	Create(book domain.Book) (domain.Book, error)
	Update(book domain.Book) (domain.Book, error)
	Delete(book domain.Book) (domain.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (domain.Book, error) {
	var book domain.Book
	err := r.db.Where("id = ?", ID).First(&book).Error

	return book, err
}

func (r *repository) Create(book domain.Book) (domain.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book domain.Book) (domain.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) Delete(book domain.Book) (domain.Book, error) {
	err := r.db.Delete(book).Error

	return book, err
}
