package book

import (
	"gorm.io/gorm"
)

func BookRegistry(db *gorm.DB) Service {
	bookRepository := NewRepository(db)
	bookService := NewService(bookRepository)

	return bookService
}
