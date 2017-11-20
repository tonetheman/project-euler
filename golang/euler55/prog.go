package main

import (
	"fmt"
	"strconv"
)

const debug = true

func isPal(n int64) bool {

	// this is where we figure out the lead digit
	var divisor int64 = 1
	for n/divisor >= 10 {
		divisor *= 10
	}

	//fmt.Println("isPal divisor", divisor)

	for n != 0 {
		leading := n / divisor
		trailing := n % 10
		//fmt.Println("leading,trailing", leading, trailing)

		if leading != trailing {
			return false
		}

		n = (n % divisor) / 10

		divisor = divisor / 100
	}

	return true
}

// i did not write this function
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isLychrel(n int64) int {
	if isPal(n) {
		return 1
	}
	if debug {
		fmt.Println("input", n)
	}
	count := 0
	for count < 50 {
		ns := strconv.FormatInt(int64(n), 10)
		nsr := Reverse(ns)
		r, _ := strconv.ParseInt(nsr, 10, 32)
		newN := int(r) + n
		newIsPal := isPal(newN)
		count++ // this iteration has been done

		if debug {
			fmt.Println(newN, newIsPal)

		}
		if newIsPal {
			break
		}

		n = newN
	}
	if debug {
		fmt.Println("total iterations", count)
	}
	return count
}

func testit() {
	for i := 10; i < 100; i++ {
		if isLychrel(i) > 30 {
			fmt.Println(i)
		}
	}
}

func main() {

	fmt.Println(isLychrel(98))
	//fmt.Println(isLychrel(545))

}
