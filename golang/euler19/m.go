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
	// this is the years
	for i := 1901; i < 2001; i++ {
		// this is the days
		for j := 1; j < 32; j++ {
			d := time.Date(i, time.January, j, 0, 0, 0, 0, locl)
			fmt.Println(d, d.Day(), d.Weekday())
			if d.Weekday() == time.Sunday {
				sundayCount++
			}
		}
	}
	fmt.Println("sunday count", sundayCount)
}
