package main

type BookShelfIterator struct {
	bookShelf BookShelf
	index     int
}

func (bsi BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (bsi BookShelfIterator) Next() Book {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}
