package builder

import (
	"reflect"
	"testing"
)

func TestNewTextBuilder(t *testing.T) {
	n := NewTextBuilder()
	if reflect.TypeOf(n).String() != "*builder.TextBuilder" {
		t.Error("NewTextBuilder failed.")
	}
}

func TestNewHTMLBuilder(t *testing.T) {
	n := NewHTMLBuilder()
	if reflect.TypeOf(n).String() != "*builder.HTMLBuilder" {
		t.Error("NewHTMLBuilder failed.")
	}
}

func TestNewDirector(t *testing.T) {
	h := NewHTMLBuilder()
	n := NewDirector(h)
	if reflect.TypeOf(n).String() != "*builder.Director" {
		t.Error("NewDirector failed.")
	}
}

func TestHTMLBuilderConstruct(t *testing.T) {
	h := NewHTMLBuilder()
	d := NewDirector(h)
	d.Construct()
}

func TestTextBuilderConstruct(t *testing.T) {
	tx := NewTextBuilder()
	d := NewDirector(tx)
	d.Construct()
}
