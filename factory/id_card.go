package factory

import (
	"fmt"
)

// IDCard implement Product. IDCard can use.
type IDCard struct {
	owner string
}

// NewIDCard constructor
func NewIDCard(owner string) *IDCard {
	return &IDCard{owner}
}

// Use is Product's function
func (i IDCard) Use() {
	fmt.Println(i.owner + "のカードを作ります")
}

// GetOwner return owner string
func (i IDCard) GetOwner() string {
	return i.owner
}
