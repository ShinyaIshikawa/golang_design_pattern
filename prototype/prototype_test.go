package prototype

import (
	"testing"
)

func TestUnderlinePenClone(t *testing.T) {
	m := NewManager()
	upen := NewUnderlinePen("~")
	m.Register("strong message", upen)
	p := m.Create("strong message")
	if upen.ulchar == "" || p.(UnderlinePen).ulchar == "" {
		t.Error("fail clone UnderlinePen instance.")
	}
	if upen.ulchar != p.(UnderlinePen).ulchar {
		t.Error("fail clone UnderlinePen instance.")
	}
}

func TestMessageBoxClone(t *testing.T) {
	m := NewManager()
	mbox := NewMessageBox("~")
	m.Register("strong message", mbox)
	p := m.Create("strong message")
	if mbox.decochar == "" || p.(MessageBox).decochar == "" {
		t.Error("fail clone UnderlinePen instance.")
	}
	if mbox.decochar != p.(MessageBox).decochar {
		t.Error("fail clone MessageBox instance.")
	}
}

func TestUnderlinePenUse(t *testing.T) {
	m := NewManager()
	upen := NewUnderlinePen("~")
	m.Register("strong message", upen)
	p := m.Create("strong message")
	p.Use("Hello, world")
}

func TestMessageBoxUse(t *testing.T) {
	m := NewManager()
	mbox := NewMessageBox("*")
	m.Register("strong message", mbox)
	p := m.Create("strong message")
	p.Use("Hello, world")
}

func TestCloneInstanceNotFound(t *testing.T) {
	m := NewManager()
	upen := NewUnderlinePen("~")
	m.Register("strong message", upen)
	p := m.Create("strong message hoge")
	if p != nil {
		t.Error("error Manager Crate function.")
	}
}

func TestMultipleExecute(t *testing.T) {
	m := NewManager()
	upen := NewUnderlinePen("~")
	mbox := NewMessageBox("*")
	m.Register("strong message", upen)
	m.Register("strong message", mbox)
	p1 := m.Create("strong message")
	p1.Use("Hello, world")
	p2 := m.Create("strong message")
	p2.Use("Hello, world")
}
