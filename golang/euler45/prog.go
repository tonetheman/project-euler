package main

import "fmt"

const count = 0

// thought here is to keep a running copy of the current target
// tri number
// then check the two other types of numbers for matches
// until they go past the current target
// I think that once I get a fail then the number that was last checked
// will need to be backed down an index value to be careful

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
		fmt.Println("TARGET: looking for this tri number", computedTri)

		// int the computed penta number
		foundPent := false
		computedPent := 0
		for computedPent < computedTri {
			computedPent = pent(pentIndex)
			if computedPent == computedTri {
				fmt.Println("FOUND ONE")
				foundPent = true
				break
			}
			fmt.Println("computed penta:", computedPent)
			pentIndex++
		}

		counter++
		if counter > 1 {
			break
		}

		if !foundPent {
			fmt.Println("did not find a matching penta jumping back to top for new target")
			pentIndex-- // doing this to be careful
			continue
		}
		// now check hexaIndex
		computedHexa := 0
		foundHexa := false
		for computedHexa < computedTri {
			computedHexa = hexa(hexaIndex)
			if computedHexa == computedTri {
				fmt.Println("FOUND ONE!")
				foundHexa = true
			}
			fmt.Println("computed hexa", computedHexa)
			hexaIndex++
		}

		if !foundHexa {
			fmt.Println("no matching hexa found")
			hexaIndex--
			continue
		}
		if foundHexa {
			fmt.Println("TOTALL DONE")
			break
		}
	}
}
