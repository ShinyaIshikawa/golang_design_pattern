package template

import (
	"fmt"
)

// StringDisplay implement Print interface.
type StringDisplay struct {
	string string
	width  int
}

// NewStringDisplay constructor.
func NewStringDisplay(string string, width int) *StringDisplay {
	//new
	sd := new(StringDisplay)
	sd.string = string
	sd.width = width
	return sd
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
