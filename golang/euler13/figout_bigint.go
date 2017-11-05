package main

import "math/big"
import "fmt"


func main() {
  fmt.Println("hi")
  var a big.Int
  a.SetInt64(0)
  fmt.Println("just println", a)
  fmt.Println("call string", a.String())

}
