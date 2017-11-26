package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const filename string = "poker.txt"
const (
	unknown       = 0
	royalFlush    = 20
	straightFlush = 18
	fourOfKind    = 17
	fullHouse     = 16
	flush         = 15
	straight      = 14
	threeOfKind   = 13
	twoPair       = 12
	onePair       = 11
)

// Card card struct here
type Card struct {
	rank int
	suit int
}

// Hand struct here
type Hand [5]Card

func (h *Hand) Len() int {
	return len(h)
}
func (h *Hand) Swap(i, j int) {
	h[i].rank, h[j].rank = h[j].rank, h[i].rank
	h[i].suit, h[j].suit = h[j].suit, h[i].suit
}
func (h *Hand) Less(i, j int) bool {
	return h[i].rank < h[j].rank
}

const (
	ace   = 1
	ten   = 10
	jack  = 11
	queen = 12
	king  = 13
)

func translateRank(s string) int {
	var c = s[0]
	if c == 'A' {
		return ace
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
		return jack
	}
	if c == 'Q' {
		return queen
	}
	if c == 'K' {
		return king
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
	cards[0] = newCard(h[0])
	cards[1] = newCard(h[1])
	cards[2] = newCard(h[2])
	cards[3] = newCard(h[3])
	cards[4] = newCard(h[4])
	return cards
}

func checkFlush(v Hand) bool {
	return v[0].suit == v[1].suit &&
		v[1].suit == v[2].suit &&
		v[2].suit == v[3].suit &&
		v[3].suit == v[4].suit
}

// ScoreHand expects a sorted/transormed hand
func ScoreHand(h []string) int {

	v := newHand(h)
	//fmt.Println("newHand", v)

	sort.Sort(&v)
	//fmt.Println("sorted", v)

	isFlush := checkFlush(v)

	if isFlush {
		if v[0].rank == ace &&
			v[1].rank == ten &&
			v[2].rank == jack &&
			v[3].rank == queen &&
			v[4].rank == king {
			return royalFlush
		}
	}

	// STRAIGHT_FLUSH
	if isFlush {
		if v[0].rank == v[1].rank-1 &&
			v[1].rank == v[2].rank-1 &&
			v[2].rank == v[3].rank-1 &&
			v[3].rank == v[4].rank-1 {
			return straightFlush
		}
	}
	// NOTE: do not need the ACE leading case
	// since that is a royal flush

	// 4 of a kind left
	if v[0].rank == v[1].rank &&
		v[1].rank == v[2].rank &&
		v[2].rank == v[3].rank &&
		v[3].rank != v[4].rank {
		return fourOfKind
	}
	// 4 of a kind right
	if v[1].rank == v[2].rank &&
		v[2].rank == v[3].rank &&
		v[3].rank == v[4].rank &&
		v[0].rank != v[1].rank {
		return fourOfKind
	}

	// full house
	if v[0].rank == v[1].rank &&
		v[1].rank == v[2].rank &&
		v[3].rank == v[4].rank {
		return fullHouse
	}
	if v[0].rank == v[1].rank &&
		v[2].rank == v[3].rank &&
		v[3].rank == v[4].rank {
		return fullHouse
	}

	// flush
	if isFlush {
		return flush
	}

	// straight
	if v[0].rank == v[1].rank-1 &&
		v[1].rank == v[2].rank-1 &&
		v[2].rank == v[3].rank-1 &&
		v[3].rank == v[4].rank-1 {
		return straight
	}
	// checking for the case where the ACE leads
	if v[0].rank == ace &&
		v[1].rank == v[2].rank-1 &&
		v[2].rank == v[3].rank-1 &&
		v[3].rank == v[4].rank-1 {
		return straight
	}

	// 3 of kind
	if v[0].rank == v[1].rank &&
		v[1].rank == v[2].rank &&
		v[2].rank != v[3].rank &&
		v[2].rank != v[4].rank {
		return threeOfKind
	}
	if v[4].rank == v[3].rank &&
		v[3].rank == v[2].rank &&
		v[2].rank != v[1].rank &&
		v[2].rank != v[0].rank {
		return threeOfKind
	}

	// two pair
	// 3 cases here 2 x 2
	if v[0].rank == v[1].rank &&
		v[2].rank != v[1].rank &&
		v[2].rank != v[3].rank &&
		v[3].rank == v[4].rank {
		return twoPair
	}
	// 2 2 x
	if v[0].rank == v[1].rank &&
		v[2].rank == v[3].rank &&
		v[4].rank != v[3].rank &&
		v[4].rank != v[1].rank {
		return twoPair
	}
	// x 2 2
	if v[0].rank != v[1].rank &&
		v[0].rank != v[3].rank &&
		v[1].rank == v[2].rank &&
		v[3].rank == v[4].rank {
		return twoPair
	}

	// one pair
	if v[0].rank == v[1].rank {
		return onePair
	}
	if v[1].rank == v[2].rank {
		return onePair
	}
	if v[2].rank == v[3].rank {
		return onePair
	}
	if v[3].rank == v[4].rank {
		return onePair
	}

	// high card
	// TODO: high card!

	return unknown
}

func getHiCard(h []string) int {
	v := newHand(h)
	sort.Sort(&v)
	var hiCard int
	if v[0].rank == ace {
		hiCard = 99
	} else {
		hiCard = v[4].rank
	}
	return hiCard
}

// HandleLine reads in a line as a string
func HandleLine(inputData string) ([]string, []string) {
	d := strings.Split(inputData, " ")
	for i := 0; i < len(d); i++ {
		d[i] = strings.TrimSpace(d[i])
	}
	h0 := d[0:5]
	h1 := d[5:10]
	return h0, h1
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
	playerOneWinCount := 0

	for i := 0; i < len(data); i++ {
		d := strings.TrimSpace(data[i])
		if len(d) > 0 {
			totalCount++
			h0, h1 := HandleLine(d)
			h0Score := ScoreHand(h0)
			h1Score := ScoreHand(h1)

			if h0Score == unknown && h1Score == unknown {
				unknownCount++

				// need to figure out high card
				hiCard0 := getHiCard(h0)
				hiCard1 := getHiCard(h1)
				if hiCard0 > hiCard1 {
					playerOneWinCount++
				}
				if hiCard0 == hiCard1 {
					fmt.Println("HI CARD TIE!!!")
				}
			} else if h0Score > h1Score {
				playerOneWinCount++
			}

		}
	}
	fmt.Println("unknown hands", unknownCount)
	fmt.Println("playerOne win count", playerOneWinCount)
}
