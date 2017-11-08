package main

import "fmt"
import "math/big"

const COUNT int = 20000
func main() {

  var f [COUNT]big.Int
  f[0].SetInt64(1)
  f[1].SetInt64(1)
  for i := 2; i<COUNT; i++ {
    f[i].Add(&f[i-1],&f[i-2])
  }

  //fmt.Println(f[len(f)-1].String())

  for i:=0;i<COUNT;i++ {
    if len(f[i].String())== 1000 {
      fmt.Println(i+1)
    }
  }
}
