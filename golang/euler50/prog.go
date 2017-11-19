package main

import (
	"fmt"
	"math"
)

func checkPrime(i int, a []int) bool {
	if i%2 == 0 {
		return false
	}
	top := int(math.Sqrt(float64(i))) + 1
	for j := 3; j < top; j += 2 {
		if i%j == 0 {
			return true
		}
	}
	return false
}

func createPrimeSlice(limit int) []int {
	var primes []int

	primes = append(primes, 2)
	primes = append(primes, 3)

	for i := 5; i < limit; i += 2 {
		if checkPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	fmt.Println("hi")

	primes := createPrimeSlice(1000)
	fmt.Println(primes)

}
