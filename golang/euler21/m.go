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
  // removed looking
  // for proper Divisors aka less than n
  //res[index] = n
  //index++
  _res := make([]int,index)
  copy(_res,res[0:index])
  return _res
}

func Sum(a []int) (total int) {
  total = 0
  for i:=0;i<len(a);i++ {
    total += a[i]
  }
  return
}

var mem map[int][]int

func main() {
  mem = make(map[int][]int)

  a := Divisors(220)
  mem[220] = a
  fmt.Println(mem)
  fmt.Println(a)
  total := Sum(a)
  fmt.Println(total)
}
