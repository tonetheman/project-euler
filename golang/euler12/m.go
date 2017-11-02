package main

import "fmt"

const tri1 int64 = 1
var current_tri int64 = 1

func factors(n int64) int {
  var i int64
  var facs [500]int64
  var counter int = 2
  for i=2;i<(n/2)+1;i++ {
    if n%i == 0 {
      facs[counter]=i
      counter = counter + 1
    }
  }
  return counter
}

func main() {

  var i int64
  for i = 2;i<100000;i++ {
    current_tri = current_tri + i
    counter := factors(current_tri)
    if counter >400 {
      fmt.Println("current triangle number", current_tri, counter)
    }
  }
}
