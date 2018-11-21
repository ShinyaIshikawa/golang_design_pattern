package main

type BookShelf struct {
	books []Book
	last  int
}

func (bs BookShelf) iterator() Iterator {
	return
}
