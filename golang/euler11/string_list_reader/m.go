package string_list_reader 

import (
	"fmt"
	"strings"
	"strconv"
)

const ROWSIZE = 5
const COLSIZE = 5
const s string = `1 1 1 1 1
2 2 2 2 2
3 3 3 3 3
4 4 4 4 4
5 5 5 5 5`

func load_data_var(s string, rowsize int, colsize int) (res [][]int) {
	var output_res [][]int = make([rowsize][colsize]int)

	return output_res
}

func load_data() (res [5][5]int) {
	data := strings.Fields(s)
	var idx,jdx = 0,0
	for i:=0;i<len(data);i++ {
		tmpi,err := strconv.ParseInt(data[i],10,8)
		if err != nil {
			fmt.Println("ERR",err)
		}
		// row col here
		// cast is needed to handle the int64 from ParseInt
		res[idx][jdx] = int(tmpi)
		jdx++
		if jdx>COLSIZE-1 {
			jdx = 0
			idx++
		}
	}
	return
}

