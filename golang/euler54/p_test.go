package main

import "strings"
import "testing"
import "fmt"

func TestScoreFullHouse1(t * testing.T) {
  s := "2H 2D 4C 4D 4S"
  f := strings.Fields(s)
  fmt.Println("FULL_HOUSE testing",f)
  v := ScoreHand(f)
  if v != FULL_HOUSE {
    t.Error("expected FULL_HOUSE(3) got",v)
  }
}

func TestScoreFullHouse2(t * testing.T) {
  s := "2H 4C 4D 4S 2D"
  f := strings.Fields(s)
  fmt.Println("FULL_HOUSE testing",f)
  v := ScoreHand(f)
  if v != FULL_HOUSE {
    t.Error("expected FULL_HOUSE(3) got",v)
  }
}

func TestScoreFullHouse3(t * testing.T) {
  s := "8H 4C 4D 4S 8D"
  f := strings.Fields(s)
  fmt.Println("FULL_HOUSE testing",f)
  v := ScoreHand(f)
  if v != FULL_HOUSE {
    t.Error("expected FULL_HOUSE(3) got",v)
  }
}

func TestFourOfKind(t * testing.T) {
  s := "5H 5D 5S 5C 4D"
  f := strings.Fields(s)
  fmt.Println("FOUR OF KIND TESTING",f)
  v := ScoreHand(f)
  if v != FOUR_OF_KIND {
    t.Error("expect four of kind got",v)
  }
}

func TestThreeOfKind(t * testing.T) {
  fmt.Println("TestThreeOfKind")
  s := "2H 2C 2S 5D TC"
  f := strings.Fields(s)
  v := ScoreHand(f)
  if v!= THREE_OF_KIND {
    t.Error("failed 3 of kind",v)
  }
}
