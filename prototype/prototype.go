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
	if m.m == nil {
		m.m = map[string]Product{}
	}
	m.m[name] = p
}

// Create call createClone function
func (m *Manager) Create(name string) Product {
	if p, ok := m.m[name]; ok {
		return p.createClone()
	}
	return nil
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

// Use print strings to standard output.
func (m MessageBox) Use(s string) {
	len := len(s)
	for i := 0; i < (len + 4); i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println("")
	fmt.Println(m.decochar + " " + s + " " + m.decochar)
	for i := 0; i < (len + 4); i++ {
		fmt.Print(m.decochar)
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

// Use print strings to standard output.
func (u UnderlinePen) Use(s string) {
	len := len(s)
	fmt.Println("¥" + s + "¥")
	fmt.Print(" ")
	for i := 0; i < len; i++ {
		fmt.Print(u.ulchar)
	}
	fmt.Println("")
}

func (u UnderlinePen) createClone() Product {
	return u
}
