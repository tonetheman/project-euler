package main

import "fmt"

func main() {

  // TONY: all 3 of these are valid!
  type Grid [][]int
  //type Grid [20][]int
  //type Grid [20][20]int

  var a Grid = Grid{
    {8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
   {49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
 }

 fmt.Println(a)
 for i:=0;i<len(a);i++ {
   fmt.Println(len(a[i]))
 }
}
