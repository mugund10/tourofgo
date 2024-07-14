package main

import (
	"fmt"
)

// Book represents a book in the library
type Book struct {
	Title  string
	Author string
	Year   int
}

// NewBook creates a new book with the given parameters
func NewBook(title, author string, year int) Book {
	return Book{title, author, year}
}

// UpdateYear updates the year of the book
func UpdateYear(book Book, newYear int) Book {
	book.Year = newYear
	return book
}

// Example usage
func main() {
	book := NewBook("1984", "George Orwell", 1949)
	fmt.Printf("Original Book: %+v\n", book)

	updatedBook := UpdateYear(book, 1950)
	fmt.Printf("Updated Book: %+v\n", updatedBook)
}
