package strategy

import (
	"math/rand"
	"time"
)

type strategy interface {
	nextHand() Hand
	study(win bool)
}

type WinningStrategy struct {
	won    bool
	random int
}

func NewWinningStrategy() WinningStrategy {
	rand.Seed(time.Now().UnixNano())
	random := rand.Int()
	return WinningStrategy{won: false, random: random}
}
