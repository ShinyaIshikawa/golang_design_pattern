package bridge

import (
	"testing"
)

func TestNewStiringDisplay(t *testing.T) {
	imp := NewStringDisplayImpl("yokohama")
	gotStr := imp.str
	wantStr := "yokohama"
	if gotStr != wantStr {
		t.Errorf("got %s, want %s", gotStr, wantStr)
	}
	gotWidth := imp.width
	wantWidth := 8
	if gotWidth != wantWidth {
		t.Errorf("got %d, want %d", gotWidth, wantWidth)
	}
}

func ExampleDisplay() {
	imp := NewStringDisplayImpl("yokohama")
	d := NewDisplay(imp)
	d.Open()
	// Output: +--------+
}
