package domain

import (
	"encoding/json"
	"time"
)

type Book struct {
	ID          int
	Title       string `json:"title" `
	Description string `json:"description" `
	Price       int    `json:"price" `
	Rating      int    `json:"rating"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookRequest struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

type BookResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}
