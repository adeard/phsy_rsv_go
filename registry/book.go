package registry

import (
	"phsy_rsv_go/modules/book"

	"gorm.io/gorm"
)

func BookRegistry(db *gorm.DB) book.Service {
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	return bookService
}
