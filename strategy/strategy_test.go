package strategy

import (
	"strconv"
	"testing"
)

var handtests = []struct {
	in  string
	out int
}{
	{"0", 0},
	{"1", 1},
	{"2", 2},
}

func TestGetHand(t *testing.T) {
	for _, tt := range handtests {
		t.Run(tt.in, func(t *testing.T) {
			i, err := strconv.Atoi(tt.in)
			if err != nil {
				t.Errorf("fail to conbert string parameter to integer.")
			}
			s := GetHand(i)
			if s.handvalue != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

var compareStrong = []struct {
	in1 Hand
	in2 Hand
	out bool
}{
	{Hand{GUU}, Hand{GUU}, false},
	{Hand{GUU}, Hand{CHO}, true},
	{Hand{GUU}, Hand{PAA}, false},
	{Hand{CHO}, Hand{PAA}, true},
	{Hand{CHO}, Hand{CHO}, false},
	{Hand{PAA}, Hand{PAA}, false},
}

func TestIsStrongerThan(t *testing.T) {
	for _, tt := range compareStrong {
		got := tt.in1.IsStrongerThan(tt.in2)
		if got != tt.out {
			t.Errorf("in1 %v in2 %v got %v, want %v", name[tt.in1.handvalue], name[tt.in2.handvalue], got, tt.out)
		}
	}
}

var compareWeak = []struct {
	in1 Hand
	in2 Hand
	out bool
}{
	{Hand{GUU}, Hand{GUU}, false},
	{Hand{GUU}, Hand{CHO}, false},
	{Hand{GUU}, Hand{PAA}, true},
	{Hand{CHO}, Hand{PAA}, false},
	{Hand{CHO}, Hand{CHO}, false},
	{Hand{PAA}, Hand{PAA}, false},
}

func TestIsWeakerThan(t *testing.T) {
	for _, tt := range compareWeak {
		got := tt.in1.IsWeakerThan(tt.in2)
		if got != tt.out {
			t.Errorf("in1 %v in2 %v got %v, want %v", name[tt.in1.handvalue], name[tt.in2.handvalue], got, tt.out)
		}
	}
}

var compareEqual = []struct {
	in1 Hand
	in2 Hand
	out int
}{
	{Hand{GUU}, Hand{GUU}, 0},
	{Hand{CHO}, Hand{CHO}, 0},
	{Hand{PAA}, Hand{PAA}, 0},
}

func TestFightEqual(t *testing.T) {
	for _, tt := range compareEqual {
		got := tt.in1.Fight(tt.in2)
		if got != tt.out {
			t.Errorf("in1 %v in2 %v got %v, want %v", name[tt.in1.handvalue], name[tt.in2.handvalue], got, tt.out)
		}
	}
}

func TestWinnigNextHand(t *testing.T) {
	s := NewWinningStrategy(100)
	s.won = true
	got := s.NextHand()
	want := s.NextHand()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
