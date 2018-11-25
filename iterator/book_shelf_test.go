package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	bookShelf.AppendBook(Book{Name: "hoge"})
	book := bookShelf.GetBookAt(0)
	assert.Equal(t,"hoge", book.Name)
}

func TestGetBookAt(t *testing.T) {
	book1 := Book{Name: "hoge"}
	bookShelf := BookShelf{Last: 0}
	bookShelf.AppendBook(book1)
	book2 := bookShelf.GetBookAt(0)
	assert.Equal(t,book1, book2)
}

func TestGetLength(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	bookShelf.AppendBook(Book{Name: "hoge"})
	bookShelf.AppendBook(Book{Name: "fuga"})
	bookShelf.AppendBook(Book{Name: "bar"})
	bookShelf.AppendBook(Book{Name: "foo"})
	bookShelf.GetLength()
	assert.Equal(t,4, bookShelf.GetLength())
}