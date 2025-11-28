package structembedding

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import "encoding/json"

// Author represents information about the book's author
type Author struct {
	// TODO: Define the Author struct fields
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Book represents information about a book
type Book struct {
	// TODO: Define the Book struct fields, embedding the Author struct
	Title  string `json:"title"`
	Author Author `json:"author"`
	Pages  int    `json:"pages"`
	ISBN   string `json:"ISBN"`
}

// Article represents information about a article
type Article struct {
	// TODO: Define the Article struct fields, embedding the Author struct
	Title   string `json:"title"`
	Author  Author `json:"author"`
	Journal string `json:"journal"`
	Year    int    `json:"year"`
}

// ParseBook parses the given JSON data into a Book struct
func ParseBook(jsonData []byte) (Book, error) {
	var b Book
	err := json.Unmarshal(jsonData, &b)
	return b, err
}

// ParseArticle parses the given JSON data into a Article struct
func ParseArticle(jsonData []byte) (Article, error) {
	var a Article
	err := json.Unmarshal(jsonData, &a)
	return a, err
}
