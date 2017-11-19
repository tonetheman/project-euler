package main

func isPal(n int) bool {

	// this is where we figure out the lead digit
	var divisor = 1
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

func isLychrel(n int) bool {
	return false
}

func main() {

	isPal(22345)
}
