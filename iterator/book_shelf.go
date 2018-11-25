package iterator

// BookShelf has book
type BookShelf struct {
	Books []Book
	Last  int
}

// GetBookAt return Book
func (bs *BookShelf) GetBookAt(index int) Book {
	return bs.Books[index]
}

// AppendBook add Book into Book slice
func (bs *BookShelf) AppendBook(book Book) {
	bs.Books = append(bs.Books, book)
	bs.Last++
}

// GetLength return books number
func (bs *BookShelf) GetLength() int {
	return bs.Last
}

// Iterator return BookShelfIterator
func (bs BookShelf) Iterator() BookShelfIterator {
	return BookShelfIterator{bookShelf: bs, index: 0}
}
