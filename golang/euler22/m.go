package main

import (
  "fmt"
  "io/ioutil"
  "sort"
  "strings"
)

const filename string = "p022_names.txt"

func score(name string, index int) int {
  var total int = 0
  for i:=0;i<len(name);i++ {
    s := name[i]
    if s == 34 {
      continue
    }
    //fmt.Println(s,s-64)
    total += ((int(s)-64))
  }
  return total * index
}

func main() {
  content,err := ioutil.ReadFile(filename)
  if err!=nil {
    fmt.Println("err reading",err)
  }

  data := strings.Split(string(content),",")
  sort.Strings(data)

  //fmt.Println(data)
  total := int64(0)
  for i:=0;i<len(data);i++ {
    s := score(data[i],i+1)
    total = total + int64(s)
  }
  fmt.Println("ttoal",total)
}
