package strategy

import (
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
		return getHand(rand.Intn(3))
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
	var bet int = rand.Int()
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
	return getHand(handvalue)
}

// NextHand return Hand instance.
func (p *ProbStrategy) getSum(hv int) int {
	var sum int
	for i := 0; i < 3; i++ {
		sum += p.history[hv][i]
	}
	return sum
}

func (p *ProbStrategy) Study(win bool) {
}

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

func getHand(handvalue int) Hand {
	return hand[handvalue]
}

func (h Hand) isStrongerThan(hand Hand) bool {
	return h.fight(hand) == 1
}

func (h Hand) isWeakerThan(hand Hand) bool {
	return h.fight(hand) == -1
}

func (h Hand) fight(hand Hand) int {
	if h == hand {
		return 0
	} else if (h.handvalue+1)%3 == hand.handvalue {
		return 1
	} else {
		return -1
	}
}

// Player is context class.
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

// Win
func (p *Player) Win() {
	p.strategy.Study(true)
	p.wincount++
	p.gamecount++
}

// Lose
func (p *Player) Lose() {
	p.strategy.Study(false)
	p.losecount++
	p.gamecount++
}

// Even.
func (p *Player) Even() {
	p.gamecount++
}
