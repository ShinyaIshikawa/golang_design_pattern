package factory

import (
	"reflect"
	"testing"
)

func TestCardFactoryType(t *testing.T) {
	icf := NewIDCardFactory()
	got := reflect.TypeOf(icf).String()
	want := "*factory.IDCardFactory"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCardCreate(t *testing.T) {
	cf := NewIDCardFactory()
	id := CreateInstance("solid principle", cf)

	value, ok := id.(*IDCard)
	if ok {
		in := "solid principle"
		got := value.owner
		want := "solid principle"
		if got != want {
			t.Errorf("Foo(%s) = %s; want %s", in, got, want)
		}
	} else {
		t.Fatal("CreateInstance function does not create IDCard instance.")
	}
}

func TestCardRegist(t *testing.T) {
	cf := NewIDCardFactory()
	CreateInstance("solid principle", cf)
	owners := cf.GetOwners()[0]
	value, ok := owners.(*IDCard)
	if ok {
		in := "solid principle"
		got := value.owner
		want := "solid principle"
		if got != want {
			t.Errorf("Foo(%s) = %s; want %s", in, got, want)
		}
	} else {
		t.Fatal("CreateInstance function does not create IDCard instance.")
	}
}
