package template

import "fmt"

// Display has display.
type Display struct {
	DisplayStrategy
}

// NewDisplay is constructor.
func NewDisplay(d DisplayStrategy) *Display {
	return &Display{d}
}

// Display function is template function
func (d Display) Display() {
	d.Open()
	for i := 0; i < 5; i++ {
		d.Print()
	}
	d.Close()
}

// DisplayStrategy interface
type DisplayStrategy interface {
	Open()
	Print()
	Close()
}

// CharDisplay implement DisplayStrategy interface.
type CharDisplay struct {
	char string
}

// NewCharDisplay constructor
func NewCharDisplay(char string) *CharDisplay {
	return &CharDisplay{char: char}
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

// StringDisplay implement DisplayStrategy interface.
type StringDisplay struct {
	string string
	width  int
}

// NewStringDisplay constructor.
func NewStringDisplay(string string, width int) *StringDisplay {
	return &StringDisplay{string: string, width: width}
}

// Open id first execute function.
func (sd StringDisplay) Open() {
	sd.printLine()
}

// Print is 2nd execute function.
func (sd StringDisplay) Print() {
	fmt.Println("|" + sd.string + "|")
}

// Close is last execute function.
func (sd StringDisplay) Close() {
	sd.printLine()
}

// printLine private function
func (sd StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < 5; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
