package main

// Book struct
type Book struct {
	name string
}

// GetName return Book name
func (b Book) GetName() string {
	return b.name
}
