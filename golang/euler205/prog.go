package main

import (
	"fmt"
	"math/rand"
	"time"
)

func peter() int {
	total := 0
	const DICECOUNT = 9
	for i := 0; i < DICECOUNT; i++ {
		tmp := rand.Intn(4) + 1
		total += tmp
	}
	return total
}

func colin() int {
	total := 0
	const DICECOUNT = 6
	for i := 0; i < DICECOUNT; i++ {
		tmp := rand.Intn(6) + 1
		total += tmp
	}
	return total
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var petewins, colinwins, draw int
	const COUNT = 1000000
	for i := 0; i < COUNT; i++ {
		p := peter()
		c := colin()
		if p > c {
			petewins++
		} else if c > p {
			colinwins++
		} else {
			draw++
		}
	}
	//fmt.Println(peter(), colin())
	fmt.Println("pete", petewins, float64(petewins)/COUNT)
	fmt.Println("colin", colinwins, float64(colinwins)/COUNT)
	fmt.Println("draw", draw, float64(draw)/COUNT)
}
