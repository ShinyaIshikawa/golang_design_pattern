package iterator

// BookShelfIterator has function for iteration
type BookShelfIterator struct {
	BookShelf BookShelf
	Index     int
}

// HasNext return boolean
func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.Index < bsi.BookShelf.GetLength() {
		return true
	}
	return false
}

// Next return Book
func (bsi *BookShelfIterator) Next() Book {
	book := bsi.BookShelf.GetBookAt(bsi.Index)
	bsi.Index++
	return book
}
