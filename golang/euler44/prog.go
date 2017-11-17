package main

import "fmt"

func pent(n int) int {
	return n * (3 *n -1)/2
}

func checkPent(n int) bool {
	for i:=0;i<COUNTER;i++ {
		if nums[i] == n {
			return true
		}
	}
	return false
}

const COUNTER = 10000

var nums [COUNTER]int

func main() {
	for i :=1; i<COUNTER+1; i++ {
		nums[i-1] = pent(i)
	}

	//fmt.Println(nums)

	for i:=0;i<COUNTER;i++ {
		for j:=0;j<COUNTER;j++ {
			if i!=j {
				num1 := nums[i]
				num2 := nums[j]
				sum := num1+num2
				diff := num1-num2
				if checkPent(sum) {
					if checkPent(diff) {
						fmt.Println("Possible",num1,num2)		
					}
					
				}
			}
		}
	}
}