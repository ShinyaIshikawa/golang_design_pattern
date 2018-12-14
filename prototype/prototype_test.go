package prototype

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	m := NewManager()
	upen := NewUnderlinePen("~")
	mbox := NewMessageBox("*")
	sbox := NewMessageBox("-")
	m.register("strong message", upen)
	m.register("waring box", mbox)
	m.register("slash box", sbox)

	p1 := m.create("strong message")
	p1.use("Hello, world")
	p2 := m.create("warning box")
	p2.use("Hello, world")
	p3 := m.create("slash box")
	p3.use("Hello, world")

}
