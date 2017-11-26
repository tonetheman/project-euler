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

func SortHand(h []string) {
	sort.Strings(h[:])
}

// expects a sorted/transormed hand
func ScoreHand(h []string) int {

	v := newHand(h)
	fmt.Println("newHand", v)

	sort.Sort(&v)
	fmt.Println("sorted", v)
	return 0

	//fmt.Println("after SortHand",h)
	// royal flush
	if h[0][0] == 'A' &&
		h[1][0] == 'W' &&
		h[2][0] == 'X' &&
		h[3][0] == 'Y' &&
		h[4][0] == 'Z' {
		if h[0][1] == h[1][1] &&
			h[1][1] == h[2][1] &&
			h[2][1] == h[3][1] &&
			h[3][1] == h[4][1] {
			fmt.Println("ROYAL_FLUSH")
			return ROYAL_FLUSH
		}
	}

	// TODO: straight flush

	// 4 of a kind left
	if h[0][0] == h[1][0] &&
		h[1][0] == h[2][0] &&
		h[2][0] == h[3][0] &&
		h[3][0] != h[4][0] {
		return FOUR_OF_KIND
	}
	// 4 of a kind right
	if h[1][0] == h[2][0] &&
		h[2][0] == h[3][0] &&
		h[3][0] == h[4][0] &&
		h[0][0] != h[1][0] {
		return FOUR_OF_KIND
	}

	// full house
	if h[0][0] == h[1][0] &&
		h[1][0] == h[2][0] &&
		h[3][0] == h[4][0] {
		return FULL_HOUSE
	}
	if h[0][0] == h[1][0] &&
		h[2][0] == h[3][0] &&
		h[3][0] == h[4][0] {
		return FULL_HOUSE
	}

	// TODO: flush
	if h[0][1] == h[1][1] &&
		h[1][1] == h[2][1] &&
		h[2][1] == h[3][1] &&
		h[3][1] == h[4][1] {
		return FLUSH
	}

	// TODO: straight

	// 3 of kind
	if h[0][0] == h[1][0] &&
		h[1][0] == h[2][0] &&
		h[2][0] != h[3][0] &&
		h[2][0] != h[4][0] {
		return THREE_OF_KIND
	}
	if h[4][0] == h[3][0] &&
		h[3][0] == h[2][0] &&
		h[2][0] != h[1][0] &&
		h[2][0] != h[0][0] {
		return THREE_OF_KIND
	}

	// two pair

	// one pair

	// high card

	return UNKNOWN
}

func HandleLine(input_data string) (int, int) {
	d := strings.Split(input_data, " ")
	for i := 0; i < len(d); i++ {
		d[i] = strings.TrimSpace(d[i])
	}
	h0 := d[0:5]
	h1 := d[5:10]
	h0_score := ScoreHand(h0)
	h1_score := ScoreHand(h1)
	fmt.Println(h0, h0_score, h1, h1_score)
	return h0_score, h1_score
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
