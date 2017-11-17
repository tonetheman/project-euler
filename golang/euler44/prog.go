package main

import "fmt"

func pent(n int) int {
	return n * (3 *n -1)/2
}

func checkPentSmart(n int) bool {
	if nums[COUNTER/2] < n {
		// it is in the lower half
		tmp := nums[0:COUNTER/2]
		for i := 0;i<len(tmp);i++ {
			if tmp[i]==n {
				return true
			}
		}
	} else {
		// it is in the top half
		tmp := nums[COUNTER/2:]
		for i := 0;i<len(tmp);i++ {
			if tmp[i]==n {
				return true
			}
		}
	}
	return false
}

func checkPent(n int) bool {
	for i:=0;i<COUNTER;i++ {
		if nums[i] == n {
			return true
		}
	}
	return false
}

const COUNTER = 1000

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
				if checkPentSmart(sum) {
					if checkPentSmart(diff) {
						fmt.Println("Possible",num1,num2)		
					}
					
				}
			}
		}
	}
}