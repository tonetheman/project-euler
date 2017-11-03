package main

import "fmt"
import "math"

const tri1 int64 = 1
var current_tri int64 = 1

func factors(n int64) int {
  var i int64

  // we have 2 factors 1, n
  var count int = 2

  // now check up to sqrt
  // each factor you find gives you 2 more
  // the number you found plus the matching number
  // above the sqrt
  for i=2;float64(i)<math.Sqrt(float64(n))+1;i++ {
    if n%i == 0 {
      count = count + 2
    }
  }
  //fmt.Println(n,facs)
  return count
}

func main() {

  var i int64
  for i = 2;i<100000;i++ {
    current_tri = current_tri + i
    counter := factors(current_tri)
    if counter >500 {
      fmt.Println("current triangle number", current_tri, counter)
      return
    }
  }
}
