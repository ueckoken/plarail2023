package main

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestDBAdd(t *testing.T) {
	// mongodb add test using db.go
	godotenv.Load(".env")
	open()
	defer close()
	book := Book{
		Title:  "Test Book",
		Author: "Test Author",
	}
	err := addBook(book)
	if err != nil {
		t.Error(err)
	}

	result, err := getBook(book.Title)
	if err != nil {
		t.Error(err)
	}

	// check if book is a match
	if result.Author != book.Author && result.Title != book.Title {
		t.Error("Book title does not match")
	}

	defer func() {
		err = deleteBook(book)
		if err != nil {
			t.Error(err)
		}
	}()
}
