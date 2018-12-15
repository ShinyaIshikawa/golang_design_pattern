package prototype

import (
	"fmt"
)

// Product is Prototype interface.
// Define clone function.
type Product interface {
	createClone() Product
	Use(s string)
}

// Manager has Procuct instance.
type Manager struct {
	m map[string]Product
}

// NewManager is Manager consructor.
func NewManager() Manager {
	return Manager{}
}

// Register add Prodct to map.
func (m *Manager) Register(name string, p Product) {
	m.m = map[string]Product{}
	m.m[name] = p
}

// Create call createClone function
func (m *Manager) Create(name string) Product {
	p := m.m[name]
	return p.createClone()
}

// MessageBox has decoracte character.
// Decorator character was printed standard output.
type MessageBox struct {
	decochar string
}

// NewMessageBox is MessageBox consructor.
func NewMessageBox(s string) MessageBox {
	return MessageBox{s}
}

func (m MessageBox) Use(s string) {
	len := len(s)
	for i := 0; i < len; i++ {
		fmt.Println(m.decochar)
	}
	fmt.Println("")
	fmt.Println(m.decochar + " " + s + " " + m.decochar)
	for i := 0; i < (len + 4); i++ {
		fmt.Println(m.decochar)
	}
	fmt.Println("")
}

func (m MessageBox) createClone() Product {
	return m
}

// UnderlinePen han underline character.
// Client tell underline character.
type UnderlinePen struct {
	ulchar string
}

// NewUnderlinePen is nderlinePen consructor.
func NewUnderlinePen(s string) UnderlinePen {
	return UnderlinePen{s}
}

func (u UnderlinePen) Use(s string) {
	len := len(s)
	fmt.Println("¥" + s + "¥")
	fmt.Println("")
	for i := 0; i < len; i++ {
		fmt.Println(u.ulchar)
	}
	fmt.Println("")
}

func (u UnderlinePen) createClone() Product {
	return u
}
