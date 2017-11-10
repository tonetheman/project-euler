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
  FOUR_OF_KIND = 2
  FULL_HOUSE = 3
  THREE_OF_KIND = 4
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
  //fmt.Println("enter ScoreHand",h)
  TransformHand(h)
  //fmt.Println("after TransformHand",h)
  SortHand(h)
  //fmt.Println("after SortHand",h)
  // royal flush
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
    // 4 of a kind left
    if h[0][0] == h[1][0] &&
      h[1][0] == h[2][0] &&
      h[2][0] == h[3][0] &&
      h[3][0] != h[4][0] {
        return FOUR_OF_KIND;
      }
      // 4 of a kind right
    if h[1][0] == h[2][0] &&
      h[2][0] == h[3][0] &&
      h[3][0] == h[4][0] &&
      h[0][0] != h[1][0] {
        return FOUR_OF_KIND;
      }

      // full house
      if h[0][0] == h[1][0] &&
        h[1][0] == h[2][0] &&
        h[3][0] == h[4][0] {
          return FULL_HOUSE
        }
      if h[0][0] == h[1][0] &&
        h[2][0] == h[3][0] &&
        h[3][0] == h[4][0] {
          return FULL_HOUSE
        }
      if h[0][0] == h[1][0] &&
        h[1][0] == h[2][0] &&
        h[2][0] != h[3][0] &&
        h[2][0] != h[4][0] {
          return THREE_OF_KIND
        }
      if h[4][0] == h[3][0] &&
        h[3][0] == h[2][0] &&
        h[2][0] != h[1][0] &&
        h[2][0] != h[0][0] {
          return THREE_OF_KIND
        }
      return UNKNOWN
}


func HandleLine(input_data string) (int, int){
  d := strings.Split(input_data," ")
  for i:=0;i<len(d);i++ {
    d[i] = strings.TrimSpace(d[i])
  }
  h0 := d[0:5]
  h1 := d[5:10]
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
  data[0] = strings.TrimSpace(data[0])
  if len(data[0]) >0 {
    HandleLine(data[0])
  }
}
