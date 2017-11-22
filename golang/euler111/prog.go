package main

import "fmt"
import "math"

var primes [1000]int

func printMaxValues() {
	fmt.Println("maxint32", math.MaxInt32)
	fmt.Println("maxint64", math.MaxInt64)
	fmt.Println("maxfloat64", math.MaxFloat64)
}

// determines top of loop needed to check
// for primes efficiently
func topOfSearch(n int64) int64 {
	return int64(math.Sqrt(float64(n))) + 1
}

func isPrime(n int) bool {
	return false
}

func main() {
	fmt.Println("starting")
	printMaxValues()
	a := int64(9999999999)
	fmt.Println(topOfSearch(a))
	primes[0] = 2
	primes[1] = 3
	currentCount := 2
	for i := 5; i < int(topOfSearch(a)); i += 2 {
		if isPrime(i) {
			currentCount++
			primes[currentCount-1] = i
		}
	}
}
