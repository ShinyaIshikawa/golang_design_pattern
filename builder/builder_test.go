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
