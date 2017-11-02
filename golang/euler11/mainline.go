package main

import (
  "fmt"
  "strconv"
  "strings"
)

const s string = `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`



func load_data() [20][20]int {
  fields := strings.Fields(s)
  fmt.Println(len(fields))
  var data [20][20]int
  idx := 0
  jdx := 0
  for i:=0;i<len(fields);i++ {
		// throwing away the value for now
		// ParseInt takes a string
		// a base and a bit size
		tmpi,err := strconv.ParseInt(fields[i],10,16)
		if err != nil {
			fmt.Println("ERROR", err)
		}
    data[idx][jdx] = int(tmpi)
    idx = idx + 1
    if idx > 19 {
      idx = 0
      jdx = jdx + 1
    }
  }

    return data
}

func find_max_horizontal(data [20][20]int) {
  // search for max horizontal 4 values
  var max_h_i, max_h_ridx int
  var max_h = 0
  for ridx := 0; ridx <20;ridx ++ {
    for i:=0;i<20-4;i++ {
      var prod = data[ridx][i] * data[ridx][i+1] * data[ridx][i+2] + data[ridx][i+3]
      if prod > max_h {
        max_h = prod
        max_h_i = i
        max_h_ridx = ridx
      }
    }
  }
  fmt.Println("max prod h", max_h, "row", max_h_ridx, "col", max_h_i)
}

func find_max_vertical(data [20][20]int) {
  max_v := 0
  var max_v_cidx, max_v_i int
  for cidx :=0; cidx<20; cidx++ {
    for i:=0;i<20-4;i++ {
      var prod = data[i][cidx] * data[i+1][cidx] * data[i+2][cidx] * data[i+3][cidx]
      if prod > max_v {
        max_v = prod
        max_v_cidx = cidx
        max_v_i = i
      }
    }
  }
  fmt.Println("max v", max_v, "col", max_v_cidx, "row", max_v_i)
}

// not sure about this part
// it is yuch
func find_diag_LR_down(data [20][20]int) {
  var max_d = 0
  var row,col int
  for jcol:=0;jcol<20-4;jcol++ {
    for i:=0;i<20-4;i++ {
      var prod = data[i][jcol+0] * data[i+1][jcol+1] *
      data[i+2][jcol+2] * data[i+3][jcol+3]
      if prod > max_d {
        max_d = prod
        row = i
        col = 0
      }
    }
  }
  fmt.Println("find_diag_LR_down",max_d,row,col)
}

func find_diag_RL_down(data [20][20]int) {
  //var max_d, row, col int

  //for i:=3;i<20;i++ {
  //  var prod = data[0][i]
//  }

}

func main() {
  var data [20][20]int = load_data()

  find_max_horizontal(data)
  find_max_vertical(data)
  find_diag_LR_down(data)
  find_diag_RL_down(data)

}
