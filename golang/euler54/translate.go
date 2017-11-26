package main

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
