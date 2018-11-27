package main

import (
	"fmt"

	adp "github.com/ShinyaIshikawa/golang_design_pattern/adapter"
	itr "github.com/ShinyaIshikawa/golang_design_pattern/iterator"
)

func main() {
	executeIterator()
	executeAdapter()
}

func executeIterator() {
	bookShelf := itr.BookShelf{Last: 0}
	bookShelf.AppendBook(itr.Book{Name: "Around the World in 80 days"})
	bookShelf.AppendBook(itr.Book{Name: "Bidle"})
	bookShelf.AppendBook(itr.Book{Name: "Cinderella"})
	it := bookShelf.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}

func executeAdapter() {
	pb := adp.NewPrintBanner("Black Friday SALE!")
	pb.PrintWeak()
	pb.PrintStrong()
}
