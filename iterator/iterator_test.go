package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasNext(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	it := bookShelf.Iterator()
	assert.Equal(t, false, it.HasNext())
	bookShelf.AppendBook(Book{Name: "hoge"})
	assert.Equal(t, true, it.HasNext())
}

func TestNext(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	it := bookShelf.Iterator()
	book := Book{Name: "hoge"}
	bookShelf.AppendBook(book)
	assert.Equal(t, book, it.Next())
}

func TestAppend(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	bookShelf.AppendBook(Book{Name: "hoge"})
	book := bookShelf.GetBookAt(0)
	assert.Equal(t, "hoge", book.Name)
}

func TestGetBookAt(t *testing.T) {
	book1 := Book{Name: "hoge"}
	bookShelf := BookShelf{Last: 0}
	bookShelf.AppendBook(book1)
	book2 := bookShelf.GetBookAt(0)
	assert.Equal(t, book1, book2)
}

func TestGetLength(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	bookShelf.AppendBook(Book{Name: "hoge"})
	bookShelf.AppendBook(Book{Name: "fuga"})
	bookShelf.AppendBook(Book{Name: "bar"})
	bookShelf.AppendBook(Book{Name: "foo"})
	bookShelf.GetLength()
	assert.Equal(t, 4, bookShelf.GetLength())
}
