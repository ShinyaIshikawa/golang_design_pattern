package iterator

// Book struct
type Book struct {
	Name string
}

// GetName return Book name
func (b Book) GetName() string {
	return b.Name
}

// Aggregate interface
type Aggregate interface {
	Iterator() Iterator
}

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
func (bs *BookShelf) Iterator() *BookShelfIterator {
	return &BookShelfIterator{BookShelf: bs, Index: 0}
}

// Iterator interface
type Iterator interface {
	hasNext() bool
	next() Aggregate
}

// BookShelfIterator has function for iteration
type BookShelfIterator struct {
	BookShelf *BookShelf
	Index     int
}

// HasNext return boolean whether aggregate has next or not.
func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.Index < bsi.BookShelf.GetLength() {
		return true
	}
	return false
}

// Next function return object of aggreate.
func (bsi *BookShelfIterator) Next() Book {
	book := bsi.BookShelf.GetBookAt(bsi.Index)
	bsi.Index++
	return book
}
