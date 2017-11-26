package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const filename string = "poker.txt"
const (
	UNKNOWN        = 0
	ROYAL_FLUSH    = 1
	STRAIGHT_FLUSH = 2
	FOUR_OF_KIND   = 3
	FULL_HOUSE     = 4
	THREE_OF_KIND  = 5
	FLUSH          = 6
	STRAIGHT       = 7
	TWO_PAIR       = 8
	ONE_PAIR       = 9
)

// not sure i am using this
type Card struct {
	rank int
	suit int
}

// using this
type Hand [5]Card

func (h *Hand) Len() int {
	return len(h)
}
func (h *Hand) Swap(i, j int) {
	//fmt.Println("\t\tswap called i,j", i, j)
	h[i].rank, h[j].rank = h[j].rank, h[i].rank
	h[i].suit, h[j].suit = h[j].suit, h[i].suit
}
func (h *Hand) Less(i, j int) bool {
	//fmt.Println("\tLess is called i,j", i, j)
	return h[i].rank < h[j].rank
}

const (
	ACE   = 1
	TEN   = 10
	JACK  = 11
	QUEEN = 12
	KING  = 13
)

func translateRank(s string) int {
	var c = s[0]
	if c == 'A' {
		return ACE
	}
	if c == '2' {
		return 2
	}
	if c == '3' {
		return 3
	}
	if c == '4' {
		return 4
	}
	if c == '5' {
		return 5
	}
	if c == '6' {
		return 6
	}
	if c == '7' {
		return 7
	}
	if c == '8' {
		return 8
	}
	if c == '9' {
		return 9
	}
	if c == 'T' {
		return 10
	}
	if c == 'J' {
		return JACK
	}
	if c == 'Q' {
		return QUEEN
	}
	if c == 'K' {
		return KING
	}
	return -1
}

func translateSuit(s string) int {
	//fmt.Println("translateSuit called", s)
	if s[0] == 'C' {
		return 0
	}
	if s[0] == 'S' {
		return 1
	}
	if s[0] == 'D' {
		return 2
	}
	if s[0] == 'H' {
		return 3
	}
	return -1
}

func newCard(s string) Card {
	return Card{translateRank(string(s[0])), translateSuit(string(s[1]))}
}

func newHand(h []string) Hand {
	var cards Hand
	//fmt.Println("newHand input h", h)
	//fmt.Println("newHand cards", cards)
	cards[0] = newCard(h[0])
	cards[1] = newCard(h[1])
	cards[2] = newCard(h[2])
	cards[3] = newCard(h[3])
	cards[4] = newCard(h[4])
	return cards
}

//func SortHand(h []string) {
//	sort.Strings(h[:])
//}

func checkFlush(v Hand) bool {
	return v[0].suit == v[1].suit &&
		v[1].suit == v[2].suit &&
		v[2].suit == v[3].suit &&
		v[3].suit == v[4].suit
}

// ScoreHand expects a sorted/transormed hand
func ScoreHand(h []string) int {

	v := newHand(h)
	fmt.Println("newHand", v)

	sort.Sort(&v)
	fmt.Println("sorted", v)

	isFlush := checkFlush(v)

	if isFlush {
		if v[0].rank == ACE &&
			v[1].rank == TEN &&
			v[2].rank == JACK &&
			v[3].rank == QUEEN &&
			v[4].rank == KING {
			return ROYAL_FLUSH
		}
	}

	// STRAIGHT_FLUSH
	if isFlush {
		if v[0].rank == v[1].rank-1 &&
			v[1].rank == v[2].rank-1 &&
			v[2].rank == v[3].rank-1 &&
			v[3].rank == v[4].rank-1 {
			return STRAIGHT_FLUSH
		}
	}
	// NOTE: do not need the ACE leading case
	// since that is a royal flush

	// 4 of a kind left
	if v[0].rank == v[1].rank &&
		v[1].rank == v[2].rank &&
		v[2].rank == v[3].rank &&
		v[3].rank != v[4].rank {
		return FOUR_OF_KIND
	}
	// 4 of a kind right
	if v[1].rank == v[2].rank &&
		v[2].rank == v[3].rank &&
		v[3].rank == v[4].rank &&
		v[0].rank != v[1].rank {
		return FOUR_OF_KIND
	}

	// full house
	if v[0].rank == v[1].rank &&
		v[1].rank == v[2].rank &&
		v[3].rank == v[4].rank {
		return FULL_HOUSE
	}
	if v[0].rank == v[1].rank &&
		v[2].rank == v[3].rank &&
		v[3].rank == v[4].rank {
		return FULL_HOUSE
	}

	// flush
	if isFlush {
		return FLUSH
	}

	// straight
	if v[0].rank == v[1].rank-1 &&
		v[1].rank == v[2].rank-1 &&
		v[2].rank == v[3].rank-1 &&
		v[3].rank == v[4].rank-1 {
		return STRAIGHT
	}
	// checking for the case where the ACE leads
	if v[0].rank == ACE &&
		v[1].rank == v[2].rank-1 &&
		v[2].rank == v[3].rank-1 &&
		v[3].rank == v[4].rank-1 {
		return STRAIGHT
	}

	// 3 of kind
	if v[0].rank == v[1].rank &&
		v[1].rank == v[2].rank &&
		v[2].rank != v[3].rank &&
		v[2].rank != v[4].rank {
		return THREE_OF_KIND
	}
	if v[4].rank == v[3].rank &&
		v[3].rank == v[2].rank &&
		v[2].rank != v[1].rank &&
		v[2].rank != v[0].rank {
		return THREE_OF_KIND
	}

	// two pair
	// 3 cases here 2 x 2
	if v[0].rank == v[1].rank &&
		v[2].rank != v[1].rank &&
		v[2].rank != v[3].rank &&
		v[3].rank == v[4].rank {
		return TWO_PAIR
	}
	// 2 2 x
	if v[0].rank == v[1].rank &&
		v[2].rank == v[3].rank &&
		v[4].rank != v[3].rank &&
		v[4].rank != v[1].rank {
		return TWO_PAIR
	}
	// x 2 2
	if v[0].rank != v[1].rank &&
		v[0].rank != v[3].rank &&
		v[1].rank == v[2].rank &&
		v[3].rank == v[4].rank {
		return TWO_PAIR
	}

	// one pair
	if v[0].rank == v[1].rank {
		return ONE_PAIR
	}
	if v[1].rank == v[2].rank {
		return ONE_PAIR
	}
	if v[2].rank == v[3].rank {
		return ONE_PAIR
	}
	if v[3].rank == v[4].rank {
		return ONE_PAIR
	}

	// high card

	return UNKNOWN
}

// HandleLine reads in a line as a string
func HandleLine(inputData string) (int, int) {
	d := strings.Split(inputData, " ")
	for i := 0; i < len(d); i++ {
		d[i] = strings.TrimSpace(d[i])
	}
	h0 := d[0:5]
	h1 := d[5:10]
	h0Score := ScoreHand(h0)
	h1Score := ScoreHand(h1)
	fmt.Println(h0, h0Score, h1, h1Score)
	return h0Score, h1Score
}

func main() {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("err reading", err)
	}

	var unknownCount = 0
	var totalCount = 0

	data := strings.Split(string(content), "\n")
	fmt.Println(data[1000]) // this one is empty
	fmt.Println("len", len(data))
	data[0] = strings.TrimSpace(data[0])
	if len(data[0]) > 0 {
		totalCount++
		fmt.Println("calling HandleLine")
		h0Score, _ := HandleLine(data[0])
		if h0Score == UNKNOWN {
			unknownCount++
		}
	}
	fmt.Println("unknown hands", unknownCount)
}
