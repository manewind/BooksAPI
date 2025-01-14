package models

type Book struct {
	Title         string
	Authors       []string
	Subtitle      string
	Description   string
	AverageRating float64
	PublishedDate string
}

func CreateBooks(data struct {
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
}) []Book {
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

	return books
}
