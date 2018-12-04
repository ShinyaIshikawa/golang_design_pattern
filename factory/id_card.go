package factory

import (
	"fmt"
)

// IDCard implement prodct
type IDCard struct {
	owner string
}

// NewIDCard constructor
func NewIDCard(owner string) *IDCard {
	return &IDCard{owner}
}

// Use print message
func (i IDCard) Use() {
	fmt.Println(i.owner + "のカードを作ります")
}

// getOwner
func (i IDCard) getOwner() string {
	return i.owner
}
