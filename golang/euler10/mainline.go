package main

import (
  "fmt"
  "math"
)

const MAXLEN = 1000000
const MAXPRIMESBELOW1M = 79000

func isPrime(i int, primes []int, curlen int) bool {
  for j := 0; j<curlen; j++ {
    if primes[j]==0 {
      break
    }
    if primes[j]>i {
      break
    }
    if primes[j] > int(math.Sqrt(float64(i))) {
      break
    }
    if i%primes[j]==0 {
      return false
    }
  }
  return true
}

func main() {

  // got this number by experimenting
  // number of primes less than 1,000,000 is 78948
  var primes [MAXPRIMESBELOW1M]int
  primes[0] = 2
  primes[1] = 3

  LIMIT := MAXLEN

  index := 2

  for i:=5;i<LIMIT;i+=2 {
    if isPrime(i, primes[:], index) {
      primes[index] = i
      index++
    }
  }

  //fmt.Println(primes)
  count := 0
  for i:= 0;i<MAXPRIMESBELOW1M;i++ {
    if primes[i] != 0 {
      count++
    }
  }
  fmt.Println(count)
}
