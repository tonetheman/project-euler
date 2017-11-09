package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "sort"
)
const filename string = "poker.txt"
const (
  ROYAL_FLUSH = 0
  UNKNOWN = 1
)
// not sure i am using this
type Card struct {
  rank int
  suit int
}

// using this
type Hand struct {
  h [5]string
}

func TransformHand(h []string) {
  // change TJQK in rank
  // TO WXYZ so they will sort like we want
  for i := 0; i < 5; i++ {
    if h[i][0] == 'T' {
      h[i] = "W" + string(h[i][1])
    }
    if h[i][0] == 'J' {
      h[i] = "X" + string(h[i][1])
    }
    if h[i][0] == 'Q' {
      h[i] = "Y" + string(h[i][1])
    }
    if h[i][0] == 'K' {
      h[i] = "Z" + string(h[i][1])
    }
  }
}

func SortHand(h []string) {
  sort.Strings(h[:])
}

// expects a sorted/transormed hand
func ScoreHand(h []string) int {
  if h[0][0] == 'A' &&
    h[1][0] == 'W' &&
    h[2][0] == 'X' &&
    h[3][0] == 'Y' &&
    h[4][0] == 'Z' {
      if h[0][1] == h[1][1] &&
        h[1][1] == h[2][1] &&
        h[2][1] == h[3][1] &&
        h[3][1] == h[4][1] {
          fmt.Println("ROYAL_FLUSH")
          return ROYAL_FLUSH
        }
    }

    return UNKNOWN
}


func handle_line(input_data string) (int, int){
  d := strings.Split(input_data," ")
  for i:=0;i<len(d);i++ {
    d[i] = strings.TrimSpace(d[i])
  }
  h0 := d[0:5]
  h1 := d[5:10]
  TransformHand(h0)
  SortHand(h0)
  TransformHand(h1)
  SortHand(h1)
  h0_score := ScoreHand(h0)
  h1_score := ScoreHand(h1)
  return h0_score,h1_score
}


func main() {

  content,err := ioutil.ReadFile(filename)
  if err!=nil {
    fmt.Println("err reading",err)
  }

  data := strings.Split(string(content),"\n")
  fmt.Println(data[1000]) // this one is empty
  fmt.Println("len",len(data))
  for i:=0;i<len(data);i++ {
    data[i] = strings.TrimSpace(data[i])
    if len(data[i]) >0 {
      handle_line(data[i])
    }
  }
}
