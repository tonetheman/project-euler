package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

	// string here with embedded newlines
        s := `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00`

	//
	fmt.Println(s)

	// magic here splits on whitespace!
	d := strings.Fields(s)

	// fix data
	// need to see if there is another choice here
	// to interpret strings as int with a radix
	for i:=0;i<len(d);i++ {

		// throwing away the value for now
		// ParseInt takes a string
		// a base and a bit size
		_,err := strconv.ParseInt(d[i],10,16)
		if err != nil {
			fmt.Println("ERROR", err)
		}
		if d[i] == "00" {
			d[i] = "0"
		}
		if strings.HasPrefix(d[i],"0") && len(d[i]) == 2  {
			fmt.Println("FIX",d[i], len(d[i]), d[i][1:])
			d[i] = d[i][1:]
		} 
	}

	fmt.Println(len(d))
	fmt.Println(d)

}
