package factory

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDCardFactoryType(t *testing.T) {
	icf := NewIDCardFactory()
	assert.Equal(t, "*factory.IDCardFactory", reflect.TypeOf(icf).String())

	/*
		cdc := fm.Create("ゆうきひろし", icf)
		cdc.Use()
	*/
}
