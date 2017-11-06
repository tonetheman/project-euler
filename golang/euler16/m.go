package main

import "fmt"
import "math/big"
import "strconv"

func main () {

  var n big.Int
  var x big.Int
  x.SetInt64(2)
  var y big.Int
  y.SetInt64(1000)
  n.Exp(&x,&y,nil)
  fmt.Println(n.String())

  total := 0
  s := n.String()
  for i:=0;i<len(s);i++ {
    v,err := strconv.Atoi(string(s[i]))
    if err!=nil {
      fmt.Println("ERR",err)
    }
    total += v
  }
  fmt.Println(total)


}
