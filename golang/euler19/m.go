package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.January)
	locl, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
	}
	sundayCount := 0
	// starting point
	d := time.Date(1901, time.January, 1, 0, 0, 0, 0, locl)
	// ending point
	topEnd := time.Date(2001, time.January, 1, 0, 0, 0, 0, locl)

	for d.Before(topEnd) {

		fmt.Println(d, d.Day(), d.Weekday())
		if d.Weekday() == time.Sunday && d.Day() == 1 {
			sundayCount++
		}

		d = d.Add(time.Duration(24 * time.Hour))
	}
	fmt.Println("sunday count", sundayCount)
}
