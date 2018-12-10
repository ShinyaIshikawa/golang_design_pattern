package factory

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDCardFactoryType(t *testing.T) {
	icf := NewIDCardFactory()
	assert.Equal(t, "*factory.IDCardFactory", reflect.TypeOf(icf).String())
}

func TestIDCardCreate(t *testing.T) {
	cf := NewIDCardFactory()
	idc := CreateInstance("solid", cf)
	if value, ok := idc.(*IDCard); ok {
		assert.Equal(t, "solid", value.owner)
		assert.NotEqual(t, "hoge", value.owner)
	} else {
		assert.Fail(t, "CreateInstance function does not create IDCard instance.")
	}
}
