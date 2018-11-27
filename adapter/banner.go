package adapter

import (
	"fmt"
)

// Banner is adaptee
type Banner struct {
	title string
}

// showWithParen show string with bracket
func (b Banner) showWithParen() {
	fmt.Println("(" + b.title + ")")
}

// showWithAster show string with asterisk
func (b Banner) showWithAster() {
	fmt.Println("*" + b.title + "*")
}
