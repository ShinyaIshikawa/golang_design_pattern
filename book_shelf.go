package main

// BookShelf has book
type BookShelf struct {
	books []Book
	last  int
}

// GetBookAt return Book
func (bs *BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

// AppendBook add Book into Book slice
func (bs *BookShelf) AppendBook(book Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

// GetLength return books number
func (bs *BookShelf) GetLength() int {
	return bs.last
}

// Iterator return BookShelfIterator
func (bs BookShelf) Iterator() BookShelfIterator {
	return BookShelfIterator{bookShelf: bs, index: 0}
}
