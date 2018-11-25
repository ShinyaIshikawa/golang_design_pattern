package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasNext(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	it := bookShelf.Iterator()
	assert.Equal(t,false, it.HasNext())
	bookShelf.AppendBook(Book{Name: "hoge"})
	assert.Equal(t,true, it.HasNext())
}

func TestNext(t *testing.T) {
	bookShelf := BookShelf{Last: 0}
	it := bookShelf.Iterator()
	book := Book{Name: "hoge"}
	bookShelf.AppendBook(book)
	assert.Equal(t,book, it.Next())
}
