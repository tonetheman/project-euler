package main

import (
	"fmt"
	"math/big"
)

const debug = false

func isPalBig(n big.Int) bool {
	var TEN big.Int
	TEN.SetInt64(10)
	var divisor big.Int
	divisor.SetInt64(1)
	//var count int = 0
	for {
		var tmp big.Int
		tmp.Div(&n, &divisor)
		//fmt.Println("tmp", tmp.String())
		res := tmp.Cmp(&TEN)
		//fmt.Println("res from cmp", res)
		if res == -1 {
			break
		}
		divisor.Mul(&divisor, &TEN)
		//count++
		//if count > 10 {
		//	break
		//}
	}
	//fmt.Println("divisor", divisor.String())

	var ZERO big.Int
	var HUNDRED big.Int
	HUNDRED.SetInt64(100)

	for {
		// this is the check n != 0
		tmp := n.Cmp(&ZERO)
		if tmp == 0 {
			break
		}

		var leading big.Int
		leading.Div(&n, &divisor)
		var trailing big.Int
		trailing.Mod(&n, &TEN)

		res := leading.Cmp(&trailing)
		if res != 0 {
			return false
		}

		n.Mod(&n, &divisor)
		n.Div(&n, &TEN)
		divisor.Div(&divisor, &HUNDRED)

	}
	return true
}

func isPal(n int64) bool {

	// this is where we figure out the lead digit
	var divisor int64 = 1
	for n/divisor >= 10 {
		divisor *= 10
	}

	//fmt.Println("isPal divisor", divisor)

	for n != 0 {
		leading := n / divisor
		trailing := n % 10
		//fmt.Println("leading,trailing", leading, trailing)

		if leading != trailing {
			return false
		}

		n = (n % divisor) / 10

		divisor = divisor / 100
	}

	return true
}

// i did not write this function
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isLychrelBig(n big.Int) int {
	//fmt.Println("isLychrelbig called", n.String())
	var savedN big.Int
	savedN.Set(&n)
	if isPalBig(n) {
		return 1
	}
	n.Set(&savedN)
	count := 0
	for count < 50 {
		//fmt.Println("Value of n at the top of loop", n.String())
		ns := n.String()
		nsr := Reverse(ns)
		//fmt.Println("\tns and nsr", ns, nsr)
		var r big.Int
		r.SetString(nsr, 10)
		//fmt.Println("\tvalue of r", r.String())
		var newN big.Int
		newN.SetInt64(0)
		newN.Add(&r, &n)
		//fmt.Println("\tvalue of newN", newN.String())
		var savedNewN big.Int
		savedNewN.Set(&newN)
		newIsPal := isPalBig(newN)
		count++
		if newIsPal {
			break
		}
		//fmt.Println("\tAbout to set n to newN now", n.String(), newN.String())
		n.Set(&savedNewN)
		//fmt.Println()
	}
	return count
}

/*
func isLychrel(n int64) int {
	if isPal(n) {
		return 1
	}
	if debug {
		fmt.Println("input", n)
	}
	count := 0
	for count < 50 {
		ns := strconv.FormatInt(int64(n), 10)
		nsr := Reverse(ns)
		r, _ := strconv.ParseInt(nsr, 10, 64)
		newN := r + n
		newIsPal := isPal(newN)
		count++ // this iteration has been done

		if debug {
			fmt.Println(newN, newIsPal)

		}
		if newIsPal {
			break
		}

		n = newN
	}
	if debug {
		fmt.Println("total iterations", count)
	}
	return count
}
*/
/*
func testit() {
	var i int64
	var count int = 0
	for i = 10; i < 10000; i++ {
		if isLychrel(i) > 30 {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("total lychrel numbers", count)
}
*/
func testitBig() {
	var count int = 0
	for i := 10; i < 10000; i++ {
		var tmp big.Int
		tmp.SetInt64(int64(i))
		if isLychrelBig(tmp) > 30 {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("total lychrel numbers", count)
}

func main() {

	//var i big.Int
	//i.SetInt64(3493)
	//fmt.Println(isLychrelBig(i))
	testitBig()
}
