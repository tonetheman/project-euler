package main

import "fmt"

const (
  MON = 1
  TUE = 2
  WED = 3
  THU = 4
  FRI = 5
  SAT = 6
  SUN = 7
)

const (
  JAN = 1
  FEB = 2
  MAR = 3
  APR = 4
  MAY = 5
  JUN = 6
  JUL = 7
  AUG = 8
  SEP = 9
  OCT = 10
  NOV = 11
  DEC = 12
)

var month_day_counts map[int]int

func IsLeapYear(n int) bool {
  var res = n%4==0

  if n%100==0 {
    res = false
    if n%400 == 0 {
      res = true
    }
  }
  return res
}

func FillMonths(n int) {
    month_day_counts[JAN] = 31
    month_day_counts[FEB] = 28
    if IsLeapYear(n) {
      month_day_counts[FEB] = 29
    }
    month_day_counts[MAR] = 31
    month_day_counts[APR] = 30
    month_day_counts[MAY] = 31
    month_day_counts[JUN] = 30
    month_day_counts[JUL] = 31
    month_day_counts[AUG] = 21
    month_day_counts[SEP] = 30
    month_day_counts[OCT] = 31
    month_day_counts[NOV] = 30
    month_day_counts[DEC] = 31
}

func main() {
  fmt.Println(1)
  month_day_counts = make(map[int]int)

  for i := 1901;i<1910;i++ {
    FillMonths(i)
    fmt.Println("Year",i,IsLeapYear(i))
  }
}
