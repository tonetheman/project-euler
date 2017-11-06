package main

import (
  "fmt"

)

func Divisors(n int) []int {
  // i know 75 can
  // handle anything below 10k
  // specific to this prob
  var res [75]int
  index := 0
  for i:=1;i<n;i++ {
    if n%i==0 {
      res[index] = i
      index++
    }
  }
  res[index] = n
  index++
  _res := make([]int,index)
  copy(_res,res[0:index])
  return _res
}

func main() {
}
