package main

import "fmt"

func main() {
  var A []int = []int{23, 56, 34, 67, 15}
  B := []int{45, 56, 79, 1, 99}

  var C []int

  for i := 0; i < len(A); i++ {
    C = append(C, add(&A[i], &B[i]))
  }

  for i := range 5 {
    C[i] += 1
  }

  for i, v := range C {
    fmt.Println(i,v)
  }
}
