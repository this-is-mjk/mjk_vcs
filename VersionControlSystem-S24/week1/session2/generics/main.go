package main

import "fmt"


// type any interface{}
// Go already defines any as interface{}

func main() {
  sum := genericAdd(2,4)
  concat := genericAdd("Hi", "hello")
  addArea := genericAdd(2.0, float64(50))

  fmt.Println(sum, concat, addArea)
}

// This will give error as go compiler does not how to add two interfaces
// But this will accept any input type
//func add(a,b any) any {
//  return a+b
//}

// int add fucn
func addInt(a,b int) int {
  return a + b
}

// string add func
func addString(s1,s2 string) string {
  return s1 + s2
}

// float add func
func addFloat64(a,b float64) float64 {
  return a + b
}

// generic Add fucn
func genericAdd[T int|string|float64] (a,b T) T {
  return a + b
}
