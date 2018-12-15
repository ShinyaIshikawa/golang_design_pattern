package prototype

import (
	"testing"
)

func TestUnderlinePenClone(t *testing.T) {
	m := NewManager()
	upen := NewUnderlinePen("~")
	m.Register("strong message", upen)
	p := m.Create("strong message")
	if upen.ulchar != p.(UnderlinePen).ulchar {
		t.Error("fail clone UnderlinePen instance.")
	}
}

func TestMessageBoxClone(t *testing.T) {
	m := NewManager()
	upen := NewMessageBox("~")
	m.Register("strong message", upen)
	p := m.Create("strong message")
	if upen.decochar != p.(MessageBox).decochar {
		t.Error("fail clone MessageBox instance.")
	}
}
