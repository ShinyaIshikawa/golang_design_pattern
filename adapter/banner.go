package adapter

import (
	"fmt"
)

// Print interface
type Print interface {
	PrintWeak()
	PrintStrong()
}

// PrintBanner is adpter of Banner.
type PrintBanner struct {
	Banner
}

// NewPrintBanner is constructer.
func NewPrintBanner(title string) *PrintBanner {
	//new
	return &PrintBanner{Banner{title: title}}
}

// PrintWeak is adapter of Banner function.
func (pb PrintBanner) PrintWeak() {
	pb.showWithParen()
}

// PrintStrong is adapter of Banner function.
func (pb PrintBanner) PrintStrong() {
	pb.showWithAster()
}

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
