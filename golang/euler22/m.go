package main

import (
  "encoding/csv"
  "fmt"
  "io/ioutil"
  "strings"
)

const filename string = "p022_names.txt"

func main() {
  content,err := ioutil.ReadFile(filename)
  if err!=nil {
    fmt.Println("err reading",err)
  }

  // now read with csv stuff
  r := csv.NewReader(strings.NewReader(string(content)))
  records, err := r.Read()
  if err!=nil {
    fmt.Println("err in csv",err)
  }

  fmt.Println(records)


}
