package main

import (
	"fmt"
	itr "github.com/ShinyaIshikawa/golang_design_pattern/iterator"
	)
func main() {
	bookShelf := itr.BookShelf{Last: 0}
	bookShelf.AppendBook(itr.Book{Name: "Around the World in 80 days"})
	bookShelf.AppendBook(itr.Book{Name: "Bidle"})
	bookShelf.AppendBook(itr.Book{Name: "Cinderella"})
	it := bookShelf.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
