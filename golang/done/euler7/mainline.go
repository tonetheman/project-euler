package main

import (
  "fmt"
)

func isPrime(i int, primes [10050]int, curlen int) bool {
  for j := 0; j<curlen; j++ {
    if i%primes[j]==0 {
      return false
    }
  }
  return true
}

func main() {
  var primes [10050]int
  primes[0] = 2
  primes[1] = 3

  LIMIT := 105000

  index := 2

  for i:=5;i<LIMIT;i+=2 {
    if isPrime(i, primes, index) {
      primes[index] = i
      index++
    }

    // TODO: this is not working
    if index>LIMIT {
      break
    }
  }

  fmt.Println(primes)

  fmt.Println(primes[10000])
}
