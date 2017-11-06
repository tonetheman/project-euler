package main

import (
  "fmt"
  "math/big"
  "strconv"
)

func facfac() string{
  var z big.Int
  z.MulRange(1,100)
  //fmt.Println(z.String())
  return z.String()
}

func main() {
  s := facfac()
  total := 0
  for i:=0;i<len(s);i++ {
    b := s[i]
    ts := string(b)
    num,err := strconv.Atoi(ts)
    if err!=nil {
      fmt.Println("ERR",err)
    }
    total += num
  }
  fmt.Println(total)
}
