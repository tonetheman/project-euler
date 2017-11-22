package main

import (
  "fmt"
  "strconv"
  )

func isP(s string) bool {
  mid := len(s) /2
  last := len(s) -1
  for i := 0;i<mid;i++ {
    if s[i] != s[last-i] {
      return false
    }
  }
  return true
}

func checknum(i, j int) bool {
  product := i * j
  s := strconv.Itoa(product)
  if isP(s) {
    fmt.Println(i,j,product,"YES")
    return true
  }
  return false
}

func main() {
  fmt.Println("hey")
  var largest int = 0
  for i := 100; i<999; i++  {
    for j := 100; j<999; j++ {
      if checknum(i,j) {
        if largest < i*j {
          largest = i * j
        }
      }
    }
  }
  fmt.Println("largest",largest)
}
