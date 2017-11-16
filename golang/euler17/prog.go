package main

import "fmt"

var simple map[int]string

// ConvertNumber takes and int and returns a
// string version of the spoken number
func ConvertNumber(n int) string {
	if n <= 20 {
		return simple[n]
	}
	if n == 1000 {
		return simple[n]
	}
	if n > 20 && n < 100 {
		// compose it
		digitOne := n % 10
		leadingDigit := n / 10 * 10
		return simple[leadingDigit] + "" + simple[digitOne]
	}
	if n >= 100 {
		msDigit := n / 100
		tens := n % 100
		var last string
		if tens <= 20 {
			last = simple[tens]
		} else {
			digitOne := tens % 10
			leadingDigit := tens / 10 * 10
			last = simple[leadingDigit] + "" + simple[digitOne]
		}
		if n%100 == 0 {
			return simple[msDigit] + "hundred"
		}
		return simple[msDigit] + "hundredand" + last

	}
	return "UNK"
}
func fillOutSimple() {
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
	simple[1000] = "onethousand"
}

func main() {
	fillOutSimple()
	total := 0
	for i := 1; i < 1001; i++ {
		ts := ConvertNumber(i)
		fmt.Println(ts, len(ts))
		total += len(ts)
	}
	fmt.Println("total", total)
}
