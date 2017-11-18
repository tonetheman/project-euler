package main

import "fmt"

func pent(n int) int {
	return n * (3*n - 1) / 2
}

const COUNTER = 10000

var nums [COUNTER]int
var numMap map[int]int

func main() {
	numMap = make(map[int]int)

	for i := 1; i < COUNTER+1; i++ {
		tmp := pent(i)
		nums[i-1] = tmp
		numMap[tmp] = 1
	}

	//fmt.Println(nums)

	//fmt.Println(numMap)

	for i := 0; i < COUNTER; i++ {
		for j := 0; j < COUNTER; j++ {
			if i != j {

				num1 := nums[i]
				num2 := nums[j]
				sum := num1 + num2
				diff := num1 - num2

				// this checks if sum is in the map
				if numMap[sum] != 0 && numMap[diff] != 0 {
					fmt.Println("found one", num1, num2, sum, diff)
					fmt.Println("num1 - num2", num1-num2)
					fmt.Println("num2-num1", num2-num1)
				}
			}
		}
	}
}
