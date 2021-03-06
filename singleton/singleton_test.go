package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	expect := GetInstance()
	actual := GetInstance()

	if expect == nil || actual == nil {
		t.Errorf(`object is nil`)
	}
	if expect != actual {
		t.Errorf(`error`)
	}
}

func TestGoroutine(t *testing.T) {
	var expect *Singleton
	var actual *Singleton
	go func() {
		expect = GetInstance()
	}()
	go func() {
		actual = GetInstance()
	}()
	for {
		if expect != nil && actual != nil {
			break
		}
	}

	if expect != actual {
		t.Errorf(`error`)
	}
}
