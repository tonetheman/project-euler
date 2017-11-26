package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const filename string = "poker.txt"
const (
	ROYAL_FLUSH   = 0
	UNKNOWN       = 1
	FOUR_OF_KIND  = 2
	FULL_HOUSE    = 3
	THREE_OF_KIND = 4
	FLUSH         = 5
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

// ScoreHand expects a sorted/transormed hand
func ScoreHand(h []string) int {

	v := newHand(h)
	fmt.Println("newHand", v)

	sort.Sort(&v)
	fmt.Println("sorted", v)

	//fmt.Println("after SortHand",h)
	// royal flush
	if v[0].rank == ACE &&
		v[1].rank == KING &&
		v[2].rank == QUEEN &&
		v[3].rank == JACK &&
		v[4].rank == TEN {
		if v[0].suit == v[1].suit &&
			v[1].suit == v[2].suit &&
			v[2].suit == v[3].suit &&
			v[3].suit == v[4].suit {
			fmt.Println("ROYAL_FLUSH")
			return ROYAL_FLUSH
		}
	}

	// TODO: straight flush

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

	// TODO: flush
	if v[0].suit == v[1].suit &&
		v[1].suit == v[2].suit &&
		v[2].suit == v[3].suit &&
		v[3].suit == v[4].suit {
		return FLUSH
	}

	// TODO: straight

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

	// one pair

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

	data := strings.Split(string(content), "\n")
	fmt.Println(data[1000]) // this one is empty
	fmt.Println("len", len(data))
	data[0] = strings.TrimSpace(data[0])
	if len(data[0]) > 0 {
		fmt.Println("calling HandleLine")
		HandleLine(data[0])
	}
}
