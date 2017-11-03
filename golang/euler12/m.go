package main

import "fmt"
import "math"

const tri1 int64 = 1
var current_tri int64 = 1

var facs [600]int64
func clear_facs() {
  for i:=0;i<500;i++ {
    facs[i] = 0
  }
}

func factors(n int64) int {
  var i int64
  clear_facs()
  facs[0] = 1
  facs[1] = n
  var counter int = 2
  for i=2;float64(i)<math.Sqrt(float64(n))+1;i++ {
    if n%i == 0 {
      facs[counter]=i
      counter = counter + 1
    }
  }
  //fmt.Println(n,facs)
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
