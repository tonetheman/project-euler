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

  /*
  a := Divisors(220)
  mem[220] = a
  fmt.Println(mem)
  fmt.Println(a)
  total := Sum(a)
  fmt.Println(total)

  _,ok := mem[100]
  if ok == false {
    fmt.Println("100 not in map")
  }
  fmt.Println(Divisors(2))
  */

  fmt.Println("saving divs now...")
  // pre compute the divs
  for i:=2;i<10000;i++ {
      mem[i] = Divisors(i)
  }
  fmt.Println("finished saving")

  for i:=2;i<10000;i++ {
    for j:=2;j<10000;j++ {
      sum1 := Sum(Divisors(i))
      sum2 := Sum(Divisors(j))
      if sum1 == j && sum2 == i {
        fmt.Println(i,j)
      }
    }
  }
}
