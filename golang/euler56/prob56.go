package main

import (
  "fmt"
  "math/big"
)

func SumDigits(n big.Int) int {

  return 0
}

func main() {
  for i:=1;i<10;i++ {
    for j:=1;j<10;j++ {
      x := big.NewInt(int64(i))
      y := big.NewInt(int64(j))
      var z big.Int
      z.Exp(x,y,nil)
      SumDigits(z)
      fmt.Println(z.String())
    }
  }
}
