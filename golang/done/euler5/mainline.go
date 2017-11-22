package main

import "fmt"

func checkall(i int) bool {
  for j := 1; j<21;j++ {
    if i%j != 0 {
      return false
    }
  }
  return true
}

func main() {
  for i:=20;i<999000000;i+=20 {
    if checkall(i) {
      fmt.Println("GOT IT",i)
      break
    }
  }
}
