package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCharDisplay(t *testing.T) {
	cd := NewCharDisplay("go")
	assert.Equal(t, "go", cd.char)
}

func TestStringDisplay(t *testing.T) {
	sd := NewStringDisplay("hoge", 100)
	assert.Equal(t, "hoge", sd.string)
}
