package main

type Book struct {
	name string
}

func (b Book) getName() string {
	return b.name
}
