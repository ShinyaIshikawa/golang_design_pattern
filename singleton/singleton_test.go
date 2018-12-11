package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	obj1 := getInstance()
	obj2 := getInstance()

	expect := obj1
	actual := obj2
	if obj1 != obj2 {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
