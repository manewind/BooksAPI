package services

import (
	dba "BooksAPI/db"
	"BooksAPI/internal/models"
	"BooksAPI/internal/services"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertSQLMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db :%v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO book").
		WithArgs("Test Book",  `{"Author One","Author Two"}`,"A Great Subtitle", "2023-01-01", "A great description.", 4.5).
		WillReturnResult(sqlmock.NewResult(1, 1))

	qe := &dba.QueryExecutor{DB: db}

	book := models.Book{
		Title:         "Test Book",
		Authors:       []string{"Author One", "Author Two"},
		Subtitle:      "A Great Subtitle",
		Description:   "A great description.",
		AverageRating: 4.5,
		PublishedDate: "2023-01-01",
	}

	err = services.InsertBooks(qe, book)
	if err != nil {
		t.Fatalf("Failed to insert mock db :%v", err)
	}


	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %v", err)
	}   
}
