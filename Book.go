package main

type Book struct {
	name string
}

func (b Book) GetName() string {
	return b.name
}
