package strategy

import (
	"fmt"
	"math/rand"
)

// Strategy interface.
type Strategy interface {
	NextHand() Hand
	Study(win bool)
}

// WinningStrategy satisfy Strategy interface.
type WinningStrategy struct {
	won      bool
	prevHand Hand
}

// NewWinningStrategy return WinningStrategy instance.
func NewWinningStrategy(seed int64) WinningStrategy {
	rand.Seed(seed)
	return WinningStrategy{}
}

// NextHand return Hand instance.
func (w *WinningStrategy) NextHand() Hand {
	if !w.won {
		return GetHand(rand.Intn(3))
	}
	return w.prevHand
}

// Study set result(win or loose).
func (w *WinningStrategy) Study(win bool) {
	w.won = win
}

// ProbStrategy satisfy Strategy intarface.
type ProbStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

// NewWProbStrategy is constructor.
func NewWProbStrategy(seed int64) ProbStrategy {
	rand.Seed(seed)
	return ProbStrategy{history: [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}}
}

// NextHand return Hand instance.
func (p *ProbStrategy) NextHand() Hand {
	var bet = rand.Intn(p.getSum(p.currentHandValue))
	var handvalue int
	if bet < p.history[p.currentHandValue][0] {
		handvalue = 0
	} else if bet < p.history[p.currentHandValue][0]+p.history[p.currentHandValue][1] {
		handvalue = 1
	} else {
		handvalue = 2
	}
	p.prevHandValue = p.currentHandValue
	p.currentHandValue = handvalue
	return GetHand(handvalue)
}

// NextHand return Hand instance.
func (p *ProbStrategy) getSum(hv int) int {
	var sum int
	for i := 0; i < 3; i++ {
		sum += p.history[hv][i]
	}
	return sum
}

// Study record result.
func (p *ProbStrategy) Study(win bool) {
	if win {
		p.history[p.prevHandValue][p.currentHandValue]++
	} else {
		p.history[p.prevHandValue][(p.currentHandValue+1)%3]++
		p.history[p.prevHandValue][(p.currentHandValue+2)%3]++
	}
}

// constract rock,paper,scissors
const (
	GUU = iota
	CHO
	PAA
)

// Hand struct represent rock,paper,scissors.
type Hand struct {
	handvalue int
}

var hand = []Hand{Hand{GUU}, Hand{CHO}, Hand{PAA}}
var name = []string{"グー", " チョキ", " パー"}

// GetHand return Hand instance.
func GetHand(handvalue int) Hand {
	return hand[handvalue]
}

// IsStrongerThan return strong or not.
func (h Hand) IsStrongerThan(hand Hand) bool {
	return h.Fight(hand) == 1
}

// IsWeakerThan return weak or not.
func (h Hand) IsWeakerThan(hand Hand) bool {
	return h.Fight(hand) == -1
}

// Fight return rock–paper–scissors result.
func (h Hand) Fight(hand Hand) int {
	if h == hand {
		return 0
	} else if (h.handvalue+1)%3 == hand.handvalue {
		return 1
	} else {
		return -1
	}
}

// String convert Player to string function.
func (h Hand) String() string {
	return name[h.handvalue]
}

// Player is context class in Strategy Pattern.
type Player struct {
	name      string
	strategy  Strategy
	wincount  int
	losecount int
	gamecount int
}

// NewPlayer is constructor.
func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{name: name, strategy: strategy}
}

// NextHand set next value of Hand.
func (p *Player) NextHand() {
	p.strategy.NextHand()
}

// Win record　win result.
func (p *Player) Win() {
	p.strategy.Study(true)
	p.wincount++
	p.gamecount++
}

// Lose record　lose result.
func (p *Player) Lose() {
	p.strategy.Study(false)
	p.losecount++
	p.gamecount++
}

// Even record　lose result.
func (p *Player) Even() {
	p.gamecount++
}

// String convert Player to string function.
func (p *Player) String() string {
	return fmt.Sprintf("[%s:%d games, %d win, %d lose]", p.name, p.gamecount, p.wincount, p.losecount)
}
