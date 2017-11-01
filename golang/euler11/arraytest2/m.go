package main

import (
	"fmt"
	"strings"
)

func main() {


        s := `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00`
	fmt.Println(s)
	d := strings.Split(s," ")

	// fix data
	for i:=0;i<len(d);i++ {
		if d[i] == "00" {
			d[i] = "0"
		}
		if strings.HasPrefix(d[i],"0") && len(d[i]) == 2  {
			//fmt.Println(d[i], len(d[i]), d[i][1:])
			d[i] = d[i][1:]
		}
	}

	fmt.Println(d)
}
