package main

import (
  "bufio"
  "fmt"
  "os"

)

const filename string = "p022_names.txt"
func main() {
  inf,err := os.Open(filename)
  if err!=nil {
    fmt.Println("err",err)
    return
  }
  defer inf.Close()

  scanner := bufio.NewScanner(inf)
  for scanner.Scan() {
    
  }
  fmt.Println(inf)
}
