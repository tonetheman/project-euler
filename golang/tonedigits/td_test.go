package tonedigits

import (
  "math/big"
  "testing"
  "fmt"
  )

func TestIt(t * testing.T) {
  var i big.Int
  i.SetInt64(4567)
  GetDigits(i)

  var j int64
  j = 78901
  val := GetDigitsInt64(j)
  fmt.Println(val)
}
