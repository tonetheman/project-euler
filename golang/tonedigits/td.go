package tonedigits

//import "fmt"
import "math/big"
import "strconv"

func GetDigits(n big.Int) []int {
  // see below for algo used
  s := n.String()
  slen := len(s)
  data := make([]int,slen)
  index := 0
  for i:=0;i<slen;i++ {
    var tmp big.Int
    var tmp1 big.Int
    tmp1.SetInt64(10)
    tmp.Mod(&n,&tmp1)
    data[index] = int(tmp.Uint64())
    index ++
    n.Div(&n,&tmp1)
  }
  return data
}

func GetDigitsInt64(n int64) []int {

  // convert to string
  s := strconv.FormatInt(n,10)

  // figure out how many digits to return
  slen := len(s)

  // create storage to hold digits
  data := make([]int,slen)
  index := 0
  for i := 0;i<slen;i++ {
    data[index] = int(n %10)
    index ++
    n = n/10
  }

  return data
}
