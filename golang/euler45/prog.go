package main

import "fmt"

const count = 0

var triIndex = 285
var pentIndex = 165
var hexaIndex = 143

func tri(n int) int {
	return n * (n + 1) / 2
}
func pent(n int) int {
	return n * (3*n - 1) / 2
}
func hexa(n int) int {
	return n * (2*n - 1)
}

func main() {
	// stupid go no while loop
	counter := 0
	for {
		triIndex++
		computedTri := tri(triIndex)
		fmt.Println("looking for this tri number", computedTri)

		// int the computed penta number
		computedPent := 0
		for computedPent < computedTri {
			computedPent = pent(pentIndex)
			if computedPent == computedTri {
				fmt.Println("FOUND ONE")
			}
			fmt.Println("computed penta:", computedPent)
			pentIndex++
		}

		counter++
		if counter > 1 {
			break
		}
	}
}
