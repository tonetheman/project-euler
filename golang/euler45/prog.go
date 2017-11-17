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

func computeNewPentIndex(currentIndex int, target int) int {
	for {
		if pent(currentIndex) < target {
			return currentIndex
		}
		currentIndex--
	}
}

func computeNewHexaIndex(currentIndex int, target int) int {
	for {
		if hexa(currentIndex) < target {
			return currentIndex
		}
		currentIndex--
	}
}

func main() {
	// stupid go no while loop
	for {
		if triIndex > 100000 {
			fmt.Println("stoping loop")
			break
		}
		fmt.Println()
		fmt.Println("top of loop")
		triIndex++
		fmt.Println("triIndex is", triIndex)
		computedTri := tri(triIndex)
		fmt.Println("TARGET: looking for this tri number", computedTri)

		// int the computed penta number
		foundPent := false
		computedPent := 0
		for computedPent < computedTri {
			computedPent = pent(pentIndex)
			if computedPent == computedTri {
				fmt.Println("FOUND ONE", computedTri, computedPent)
				foundPent = true
				break
			}
			fmt.Println("computed penta:", computedPent)
			fmt.Println("index for penta match", pentIndex)
			pentIndex++
		}

		if !foundPent {
			fmt.Println("no pentmatch, computing new index")
			pentIndex = computeNewPentIndex(pentIndex, computedTri)
			fmt.Println("pentIndex is now set to", pentIndex)
			continue
		} else {
			fmt.Println("FOUND A MATCH")
			// not needed for debugging anymore
			// move to hex check
			//break
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
			hexaIndex = computeNewHexaIndex(hexaIndex, computedTri)
			fmt.Println("new hexaindex is", hexaIndex)
			continue
		} else {
			fmt.Println("TOTALLY DONE")
			fmt.Println("indexes",triIndex,pentIndex,hexaIndex)
			fmt.Println(tri(triIndex))
			fmt.Println(pent(pentIndex))
			fmt.Println(hexa(hexaIndex))
			break
		}

	}
}
