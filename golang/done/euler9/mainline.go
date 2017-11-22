package main

import "fmt"

const MAXNUM = 1000

func main() {


  for i:=1;i<MAXNUM;i++ {
    for j := 1; j< MAXNUM; j++ {
      for k := 1; k< MAXNUM; k++ {
        var tmp = i*i + j*j
        if tmp == k*k {
          // found a triple
          // does it work
          if i+j+k == 1000 {
            fmt.Println(i,j,k)
            fmt.Println(i*j*k)
          }
        }
      }
    }

  }
}
