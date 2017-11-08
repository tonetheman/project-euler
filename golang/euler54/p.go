package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "sort"
)
const filename string = "poker.txt"

func main() {

  content,err := ioutil.ReadFile(filename)
  if err!=nil {
    fmt.Println("err reading",err)
  }

  data := strings.Split(string(content),"\n")
  fmt.Println(data[1000]) // this one is empty
  fmt.Println("len",len(data))

  fmt.Println()
  fmt.Println(data[0])
  d := strings.Split(data[0]," ")
  var hand0 [5]string
  for i:=0;i<5;i++ {
    hand0[i] = d[i]
  }
  sort.Strings(hand0[:])
  fmt.Println(hand0)

}
