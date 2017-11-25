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
type Hand struct {
	h [5]string
}

func translateRank(s string) int {
	var c = s[0]
	if c == 'A' {
		return 1
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
		return 11
	}
	if c == 'Q' {
		return 12
	}
	if c == 'K' {
		return 13
	}
	return -1
}

func translateSuit(s string) int {
	if s[1] == 'C' {
		return 0
	}
	if s[1] == 'S' {
		return 1
	}
	if s[2] == 'D' {
		return 2
	}
	if s[3] == 'H' {
		return 3
	}
	return -1
}

func newHand(h []string) [5]Card {
	var cards [5]Card
	cards[0] = Card{translateRank(h[0]), translateSuit(h[0])}
	cards[1] = Card{translateRank(h[1]), translateSuit(h[1])}
	cards[2] = Card{translateRank(h[2]), translateSuit(h[2])}
	cards[3] = Card{translateRank(h[3]), translateSuit(h[3])}
	cards[4] = Card{translateRank(h[4]), translateSuit(h[4])}
	return cards
}

func TransformHand(h []string) {
	fmt.Println("orighand", h)
	// change TJQK in rank
	// TO WXYZ so they will sort like we want
	for i := 0; i < 5; i++ {
		if h[i][0] == 'T' {
			h[i] = "W" + string(h[i][1])
		}
		if h[i][0] == 'J' {
			h[i] = "X" + string(h[i][1])
		}
		if h[i][0] == 'Q' {
			h[i] = "Y" + string(h[i][1])
		}
		if h[i][0] == 'K' {
			h[i] = "Z" + string(h[i][1])
		}
	}
}

func SortHand(h []string) {
	sort.Strings(h[:])
}

// expects a sorted/transormed hand
func ScoreHand(h []string) int {
	//fmt.Println("enter ScoreHand",h)
	TransformHand(h)
	//fmt.Println("after TransformHand",h)
	SortHand(h)
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
