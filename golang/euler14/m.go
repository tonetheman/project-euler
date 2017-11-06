package main

import "fmt"

func coll(n int) (res int) {
  res = 1
  for n != 1 {
    if n%2==0 {
      n = n/2
    } else {
      n = 3*n+1
    }
    res = res + 1
  }
  return
}



func main() {

  longest_seq := 0
  target := 0
  for i := 1;i < 1000000; i++ {
      if coll(i) > longest_seq {
        longest_seq = coll(i)
        target = i
      }
  }

  fmt.Println("target",target)
  fmt.Println("longest_seq",longest_seq)

}
