package structsinterfaces

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE


import "fmt"

// 1. Printable interface
type Printable interface {
	Info() string
	PageNum() int
}

// 2. Structs

type Book struct {
	Author string
	Title  string
	Pages  int
}

type Magazine struct {
	Title string
	Issue string
	Pages int
}

// 3. Constructors

func NewBook(Author, Title string, Pages int) Book {
	return Book{Author: Author, Title: Title, Pages: Pages}
}

func NewMagazine(Title, Issue string, Pages int) Magazine {
	return Magazine{Title: Title, Issue: Issue, Pages: Pages}
}

// 4. Implement Printable for Book and Magazine

func (b Book) Info() string {
	return fmt.Sprintf("%s, %s", b.Author, b.Title)
}

func (b Book) PageNum() int {
	return b.Pages
}

func (m Magazine) Info() string {
	return fmt.Sprintf("%s, %s", m.Title, m.Issue)
}

func (m Magazine) PageNum() int {
	return m.Pages
}
