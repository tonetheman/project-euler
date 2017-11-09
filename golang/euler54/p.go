package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "sort"
)
const filename string = "poker.txt"

// not sure i am using this
type Card struct {
  rank int
  suit int
}

// using this
type Hand struct {
  h [5]string
}

func TransformHand(h0 * Hand) {
  // change TJQK in rank
  // TO WXYZ so they will sort like we want
  for i := 0; i < 5; i++ {
    if h0.h[i][0] == 'T' {
      h0.h[i] = "W" + string(h0.h[i][1])
    }
    if h0.h[i][0] == 'J' {
      h0.h[i] = "X" + string(h0.h[i][1])
    }
    if h0.h[i][0] == 'Q' {
      h0.h[i] = "Y" + string(h0.h[i][1])
    }
    if h0.h[i][0] == 'K' {
      h0.h[i] = "Z" + string(h0.h[i][1])
    }
  }
}

func SortHand(h * Hand) {
  sort.Strings(h.h[:])
}

func main() {

  content,err := ioutil.ReadFile(filename)
  if err!=nil {
    fmt.Println("err reading",err)
  }

  data := strings.Split(string(content),"\n")
  fmt.Println(data[1000]) // this one is empty
  fmt.Println("len",len(data))

  fmt.Println()
  fmt.Println(data[0])
  d := strings.Split(data[0]," ")
  for i:=0;i<len(d);i++ {
    d[i] = strings.TrimSpace(d[i])
  }
  h0 := d[0:5]
  h1 := d[5:10]

  fmt.Println(h0)
  fmt.Println(h1)

}
