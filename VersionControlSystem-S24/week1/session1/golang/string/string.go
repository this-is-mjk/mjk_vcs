package main

import (
  "fmt"
  "strings"
)

func main() {
  var s string = "Hello"
  s2 := fmt.Sprintf("%s %s", s, s)
  fmt.Println(s2)

  x := strings.Split(s2, " ")

  for _, v := range x {
    fmt.Printf("%s ", v)
  }
}
