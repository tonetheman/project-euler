package main

import "fmt"

var simple map[int]string

func ConvertNumber(n int) string {
    if n<=20 {
      return simple[n]
    }
    if n>20 && n<100 {
      // compose it
      digit_one := n%10
      leading_digit := n/10 * 10
      return simple[leading_digit] + " " + simple[digit_one]
    }
    return "UNK"
}
func FillOutSimple() {
  simple = make(map[int]string)
  simple[1] = "one"
  simple[2] = "two"
  simple[3] = "three"
  simple[4] = "four"
  simple[5] = "five"
  simple[6] = "six"
  simple[7] = "seven"
  simple[8] = "eight"
  simple[9] = "nine"
  simple[10] = "ten"
  simple[11] = "eleven"
  simple[12] = "twelve"
  simple[13] = "thirteen"
  simple[14] = "fourteen"
  simple[15] = "fifteen"
  simple[16] = "sixteen"
  simple[17] = "seventeen"
  simple[18] = "eighteen"
  simple[19] = "nineteen"
  simple[20] = "twenty"
  simple[30] = "thirty"
  simple[40] = "forty"
  simple[50] = "fifty"
  simple[60] = "sixty"
  simple[70] = "seventy"
  simple[80] = "eighty"
  simple[90] = "ninety"
}

func main() {
  FillOutSimple()
  for i := 1;i<100;i++ {
    fmt.Println(ConvertNumber(i))

  }
}
