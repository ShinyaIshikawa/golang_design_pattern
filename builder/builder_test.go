package builder

import (
	"reflect"
	"testing"
)

func TestNewTextBuilder(t *testing.T) {
	n := NewTextBuilder()
	got := reflect.TypeOf(n).String()
	want := "*builder.TextBuilder"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNewHTMLBuilder(t *testing.T) {
	n := NewHTMLBuilder()
	got := reflect.TypeOf(n).String()
	want := "*builder.HTMLBuilder"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNewDirector(t *testing.T) {
	h := NewHTMLBuilder()
	n := NewDirector(h)
	got := reflect.TypeOf(n).String()
	want := "*builder.Director"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestMakeFile(t *testing.T) {
	h := NewHTMLBuilder()
	n := NewDirector(h)
	got := reflect.TypeOf(n).String()
	want := "*builder.Director"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
