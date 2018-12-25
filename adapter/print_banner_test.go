package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrintBanner(t *testing.T) {
	pb := NewPrintBanner("Black Friday SALE!")
	assert.Equal(t, "Black Friday SALE!", pb.title)
}
