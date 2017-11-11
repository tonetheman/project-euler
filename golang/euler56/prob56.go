package main

import (
  "fmt"
  "math/big"
  "strconv"
)

func SumDigits(n big.Int) int {
  s := n.String()
  sum := 0
  for i:=0;i<len(s);i++ {
    c := string(s[i])
    tmpi,err := strconv.Atoi(c)
    if err!=nil {
      fmt.Println("ERR",err)
    }
    sum += tmpi
  }
  return sum
}

func main() {
  max_i := -1
  max_j := -1
  max_value := -1

  for i:=1;i<100;i++ {
    for j:=1;j<100;j++ {
      x := big.NewInt(int64(i))
      y := big.NewInt(int64(j))
      var z big.Int
      z.Exp(x,y,nil)
      tots := SumDigits(z)
      if tots > max_value {
        max_i = i
        max_j = j
        max_value = tots
      }
    }
  }
  fmt.Printf("max i and max j and max value: %d %d %d\n",max_i,max_j,max_value)

}
