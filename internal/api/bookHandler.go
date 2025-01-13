package api

import (
	"BooksAPI/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url" // Import the net/url package for encoding
)

type Book struct {
	Title         string
	Authors       []string
	Subtitle      string
	Description   string
	AverageRating float64
	PublishedDate string
}

func FetchBooks(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters from the URL
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// URL encode only the 'q' parameter
	encodedQuery := url.QueryEscape(query)

	// Create the complete URL for the Books API call with the encoded query
	apiURL := fmt.Sprintf("%s?q=%s&%s", config.AppConfig.APIConfig.BooksAPIBaseURL, encodedQuery, r.URL.RawQuery[len("q="+query)+1:])

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

	// Parse the response body into a structured format
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

	// Create a list of books from the parsed data
	var books []Book
	for _, item := range data.Items {
		books = append(books, Book{
			Title:         item.VolumeInfo.Title,
			Authors:       item.VolumeInfo.Authors,
			Subtitle:      item.VolumeInfo.Subtitle,
			Description:   item.VolumeInfo.Description,
			AverageRating: item.VolumeInfo.AverageRating,
			PublishedDate: item.VolumeInfo.PublishedDate,
		})
	}

	log.Printf("Books: %+v", books)
	for _, book := range books {
		log.Printf("Title: %s, Authors: %v, Average Rating: %.2f", book.Title, book.Authors, book.AverageRating)
	}

	// Return the books as JSON in the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

