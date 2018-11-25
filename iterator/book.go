package iterator

// Book struct
type Book struct {
	Name string
}

// GetName return Book name
func (b Book) GetName() string {
	return b.Name
}
