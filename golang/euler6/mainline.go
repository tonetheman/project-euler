
package main

import "fmt"

func main() {
  var total int = 0
  var total_s int = 0

  for i:=1;i<101;i++ {
    total += i
    j := i*i
    total_s += j
  }

  sq_t := total * total
  fmt.Println("sq_t",sq_t)
  fmt.Println("total_s",total_s)
  final_s := sq_t - total_s
  fmt.Println("answer", final_s)
}
