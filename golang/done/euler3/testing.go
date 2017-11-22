package main

import (
  "fmt"
  "math"
)

func checkPrime( i int64) bool {
  if i%2==0 {
    return false
  }
  limit := int64(math.Sqrt(float64(i)))
  for j:=int64(3);j<limit;j+=2 {
    if i%j == 0 {
      return false
    }
  }
  return true
}

func main() {

  fmt.Printf("max int64 is %d\n", uint64(math.MaxUint64) )

  var target int64 = 600851475143
  //fmt.Printf("max to test is %d", uint64(math.Sqrt(target)))
  for i := int64(2); i< int64(math.Sqrt(float64(target))); i++ {
    if target% int64(i)==0 {
      if checkPrime(i) {
        fmt.Println("found one",i)
      }
    }
  }
}
