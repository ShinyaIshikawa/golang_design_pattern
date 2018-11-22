package main

import "fmt"

func main() {
	bookShelf := BookShelf{last: 0}
	bookShelf.AppendBook(Book{name: "Around the World in 80 days"})
	bookShelf.AppendBook(Book{name: "Bidle"})
	bookShelf.AppendBook(Book{name: "Cinderella"})
	it := bookShelf.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
