package bridge

import (
	"fmt"
)

// Display represent feature.
type Display struct {
	DisplayImpl
}

// Open call DisplayImpl's function rawOpen.
func (d Display) Open() {
	d.rawOpen()
}

// Print call DisplayImpl's function rawPrint.
func (d Display) Print() {
	d.rawPrint()
}

// Close call DisplayImpl's function rawClose.
func (d Display) Close() {
	d.rawClose()
}

// Display execute all Display function.
func (d Display) Display() {
	d.Open()
	d.Print()
	d.Close()
}

// CountDisplay is object composition.
// Defined one function.
type CountDisplay struct {
	Display
}

// MultipleDisplay execute Print multiple times.
func (cd CountDisplay) MultipleDisplay(cnt int) {
	cd.Open()
	for i := 0; i < cnt; i++ {
		cd.Print()
	}
}

// DisplayImpl represent implement.
type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

// StringDisplayImpl implement DispImpl.
type StringDisplayImpl struct {
	str   string
	width int
}

// NewStringDisplayImpl is constructor.
func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{str: str, width: len(str)}
}

func (sd StringDisplayImpl) rawOpen() {

}

func (sd StringDisplayImpl) rowPrint() {

}

func (sd StringDisplayImpl) rowClose() {

}

func (sd StringDisplayImpl) printLine() {
	fmt.Println("+")
	for i := 0; i < sd.width; i++ {
		fmt.Println("-")
	}
	fmt.Println("+")
}
