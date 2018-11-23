package main

// BookShelfIterator has function for iteration
type BookShelfIterator struct {
	bookShelf BookShelf
	index     int
}

// HasNext return boolean
func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	}
	return false
}

// Next return Book
func (bsi *BookShelfIterator) Next() Book {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}
