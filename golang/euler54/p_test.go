package main

import "strings"
import "testing"
import "fmt"

func TestRoyalFlush(t *testing.T) {
	func() {
		s := "JS KS TS QS AS"
		f := strings.Fields(s)
		fmt.Println("ROYAL_FLUSH testing")
		v := ScoreHand(f)
		if v != ROYAL_FLUSH {
			t.Error("expected royal flush got", v)
		}
	}()

}

func TestScoreFullHouse1(t *testing.T) {
	s := "2H 2D 4C 4D 4S"
	f := strings.Fields(s)
	fmt.Println("FULL_HOUSE testing", f)
	v := ScoreHand(f)
	if v != FULL_HOUSE {
		t.Error("expected FULL_HOUSE(3) got", v)
	}
}

func TestScoreFullHouse2(t *testing.T) {
	s := "2H 4C 4D 4S 2D"
	f := strings.Fields(s)
	fmt.Println("FULL_HOUSE testing", f)
	v := ScoreHand(f)
	if v != FULL_HOUSE {
		t.Error("expected FULL_HOUSE(3) got", v)
	}
}

func TestScoreFullHouse3(t *testing.T) {
	s := "8H 4C 4D 4S 8D"
	f := strings.Fields(s)
	fmt.Println("FULL_HOUSE testing", f)
	v := ScoreHand(f)
	if v != FULL_HOUSE {
		t.Error("expected FULL_HOUSE(3) got", v)
	}
}

func TestFourOfKind(t *testing.T) {
	s := "5H 5D 5S 5C 4D"
	f := strings.Fields(s)
	fmt.Println("FOUR OF KIND TESTING", f)
	v := ScoreHand(f)
	if v != FOUR_OF_KIND {
		t.Error("expect four of kind got", v)
	}
}

func TestThreeOfKind(t *testing.T) {
	fmt.Println("TestThreeOfKind")
	s := "2H 2C 2S 5D TC"
	f := strings.Fields(s)
	v := ScoreHand(f)
	if v != THREE_OF_KIND {
		t.Error("failed 3 of kind", v)
	}
}

func TestStraightFlush(t *testing.T) {
	r := func(s string) {
		f := strings.Fields(s)
		v := ScoreHand(f)
		if v != STRAIGHT_FLUSH {
			t.Error("failed straight flush", s, v)
		}
	}
	r("2C 3C 4C 5C 6C")
	r("AD 2D 3D 4D 5D")
	r("9S TS JS QS KS")
	//r("5H 6H 7D 8H 9H")
}

func TestStraight(t *testing.T) {
	r := func(s string) {
		f := strings.Fields(s)
		v := ScoreHand(f)
		if v != STRAIGHT {
			t.Error("failed straight", s, v)
		}
	}
	r("2C 3D 4D 5D 6S")
	r("AD 2C 3S 4H 5H")
	r("TD JD QD KD AC")
}
