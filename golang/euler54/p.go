package main

import (
  "io/ioutil"
  "fmt"
  "strings"
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

}
