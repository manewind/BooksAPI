package services

import (
	"BooksAPI/db"
	"BooksAPI/internal/models"
	"context"
)

func InsertBooks(book models.Book) error {

	query := "INSERT INTO book (title,subtitle,authors,published_date,description,average_rating)"
	ctx := context.Background()
	return nil
}
