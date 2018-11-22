package main

type BookShelf struct {
	books []Book
	last  int
}

func (bs BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

func (bs BookShelf) AppendBook(book Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs BookShelf) GetLength() int {
	return bs.last
}

func (bs BookShelf) Iterator() BookShelfIterator {
	return BookShelfIterator{bookShelf: bs, index: 0}
}
