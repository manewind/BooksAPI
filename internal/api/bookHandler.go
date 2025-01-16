package api

import (
	"BooksAPI/config"
	"BooksAPI/db"
	"BooksAPI/internal/models"
	"BooksAPI/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url" // Import the net/url package for encoding
)

func FetchBooks(w http.ResponseWriter, r *http.Request, qe *db.QueryExecutor) {
	// Parse query parameters from the URL
	query := r.URL.Query().Get("q")
	log.Println(query)

	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// URL encode the 'q' parameter
	encodedQuery := url.QueryEscape(query)

	// Prepare other query parameters (e.g., maxResults)
	maxResults := r.URL.Query().Get("maxResults")
	if maxResults == "" {
		maxResults = "10" // Default value if not provided
	}

	// Create the complete URL for the Books API call
	apiURL := fmt.Sprintf("%s?q=%s&maxResults=%s", config.AppConfig.APIConfig.BooksAPIBaseURL, encodedQuery, maxResults)

	log.Printf("Fetching data from: %s\n", apiURL)

	// Send the request to the Books API
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error: Failed to fetch data from API: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to fetch data: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("Received response from API with status code: %d\n", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Received bad response from API", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read response body: %v", err), http.StatusInternalServerError)
		return
	}

	var data struct {
		Items []struct {
			VolumeInfo struct {
				Title         string   `json:"title"`
				Authors       []string `json:"authors"`
				Subtitle      string   `json:"subtitle"`
				Description   string   `json:"description"`
				AverageRating float64  `json:"averageRating"`
				PublishedDate string   `json:"publishedDate"`
			}
		}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError)
		return
	}

	books := models.CreateBooks(data)

	for _, book := range books {
		err := services.InsertBooks(qe, book)
		if err != nil {
			log.Printf("Error while inserting book: %v", err)
			continue
		}
		log.Printf("Inserted book: %s, Authors: %v, Average Rating: %.2f", book.Title, book.Authors, book.AverageRating)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
