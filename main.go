package main

import "fmt"

func main() {
	fmt.Println("hello")
	bookShelf := BookShelf{last: 0}
	bookShelf.AppendBook(Book{name: "Around the World in 80 days"})
	bookShelf.AppendBook(Book{name: "Bidle"})
	bookShelf.AppendBook(Book{name: "Cinderella"})
	fmt.Println(bookShelf)
	it := bookShelf.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
