package services

import (
	"BooksAPI/db"
	"BooksAPI/internal/models"
	"context"
	"log"

	"github.com/lib/pq"
)

func InsertBooks(qe *db.QueryExecutor, book models.Book) error {

	authors := pq.Array(book.Authors)

	query := "INSERT INTO new_book1 (title,authors,subtitle,published_date,description,average_rating)VALUES ($1, $2, $3, $4, $5, $6)"
	ctx := context.Background()

	rowsAffected, err := qe.Exec(ctx, query, book.Title, authors, book.Subtitle, book.PublishedDate, book.Description, book.AverageRating)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return err
	}

	if rowsAffected > 0 {
		log.Println(rowsAffected)
	} else {
		log.Println("No rows affected")
	}

	return nil
}
