package template

import (
	"fmt"
)

// CharDisplay implement Print interface.
type CharDisplay struct {
	char string
}

// NewCharDisplay constructor
func NewCharDisplay(char string) *CharDisplay {
	//new
	cd := new(CharDisplay)
	cd.char = char
	return cd
}

// Open function
func (cd CharDisplay) Open() {
	fmt.Println("<<")
}

// Print function
func (cd CharDisplay) Print() {
	fmt.Println(cd.char)
}

// Close function
func (cd CharDisplay) Close() {
	fmt.Println(">>")
}
