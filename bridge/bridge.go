package bridge

import (
	"fmt"
)

type Display struct {
	DisplayImpl
}

func (d Display) Open() {
	d.rawOpen()
}

func (d Display) Print() {
	d.rawPrint()
}

func (d Display) Close() {
	d.rawClose()
}

func (d Display) Display() {
	d.Open()
	d.Print()
	d.Close()
}

type CountDisplay struct {
	Display
}

func (cd CountDisplay) MultipleDisplay(cnt int) {
	cd.Open()
	for i := 0; i < cnt; i++ {
		cd.Print()
	}
}

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

type StringDisplayImpl struct {
	str   string
	width int
}

func (sd StringDisplayImpl) rawOpen() {

}

func (sd StringDisplayImpl) rowPrint() {

}

func (sd StringDisplayImpl) rowClose() {

}

func (sd StringDisplayImpl) printLine() {
	fmt.Println()
}
