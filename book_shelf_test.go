package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	bookShelf := BookShelf{last: 0}
	bookShelf.AppendBook(Book{name: "hoge"})
	book := bookShelf.GetBookAt(0)
	assert.Equal(t,"hoge", book.name)
}

func TestGetBookAt(t *testing.T) {
	book1 := Book{name: "hoge"}
	bookShelf := BookShelf{last: 0}
	bookShelf.AppendBook(book1)
	book2 := bookShelf.GetBookAt(0)
	assert.Equal(t,book1, book2)
}

func TestGetLength(t *testing.T) {
	bookShelf := BookShelf{last: 0}
	bookShelf.AppendBook(Book{name: "hoge"})
	bookShelf.AppendBook(Book{name: "fuga"})
	bookShelf.AppendBook(Book{name: "bar"})
	bookShelf.AppendBook(Book{name: "foo"})
	bookShelf.GetLength()
	assert.Equal(t,4, bookShelf.GetLength())
}