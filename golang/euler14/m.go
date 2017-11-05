package main

import "fmt"

func handle_seq(n int) int {
  count := 1
  var res int = n
  //fmt.Println("handle_seq called",n)
  for true {
    if res == 1 {
      break
    }
    if res%2==0 {
      //EVEN
      res = res/2
      //fmt.Println("E",res)
      count++
    } else {
      // ODD
      res = 3*res + 1
      //fmt.Println("O",res)
      count++
    }
  }
  return count
}

func main() {
  const V = 1000000
  var longest int = 0
  for i:=2;i<V;i++ {
    res := handle_seq(i)
    if res > longest {
      longest = i
    }
  }
  fmt.Println("longest less than",V,longest)
}
