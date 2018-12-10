package factory

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardFactoryType(t *testing.T) {
	icf := NewIDCardFactory()
	assert.Equal(t, "*factory.IDCardFactory", reflect.TypeOf(icf).String())
}

func TestCardCreate(t *testing.T) {
	cf := NewIDCardFactory()
	id := CreateInstance("solid principle", cf)
	if value, ok := id.(*IDCard); ok {
		assert.Equal(t, "solid principle", value.owner)
		assert.NotEqual(t, "hoge", value.owner)
	} else {
		assert.Fail(t, "CreateInstance function does not create IDCard instance.")
	}
}

func TestCardRegist(t *testing.T) {
	cf := NewIDCardFactory()
	CreateInstance("solid principle", cf)
	assert.NotNil(t, cf.GetOwners)
	owners := cf.GetOwners()[0]

	if value, ok := owners.(*IDCard); ok {
		assert.Equal(t, "solid principle", value.owner)
		assert.NotEqual(t, "hoge", value.owner)
	} else {
		assert.Fail(t, "CreateInstance function does not create IDCard instance.")
	}
}
