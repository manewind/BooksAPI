package main

import (
	"BooksAPI/config"
	"BooksAPI/db"
	"BooksAPI/internal/api"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World12323123123")
	}

	http.HandleFunc("/", helloHandler)
	config.LoadConfig()

	qe := &db.QueryExecutor{}

	http.HandleFunc("/books",func(w http.ResponseWriter,r *http.Request){
		api.FetchBooks(w,r,qe)
	})
	fmt.Println(config.AppConfig.APIConfig.BooksAPIBaseURL)

	config.LoadConfig()

	cm := &db.ConnectionManager{}

	conn, err := cm.Connect(config.AppConfig.DBConfig)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	defer cm.Close()

	if err := testQuery(conn); err != nil {
		fmt.Println("Error during test query", err)
		return
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testQuery(db *sql.DB) error {
	query := "SELECT 1"
	var result int

	if err := db.QueryRow(query).Scan(&result); err != nil {
		return fmt.Errorf("test query failed: %w", err)
	}
	fmt.Println("Test query successful, result:", result)
	return nil
}
