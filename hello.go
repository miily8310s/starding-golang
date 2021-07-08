package main

import (
	"fmt"
)

func div(a, b int) (int, int) {
  q := a / b
  r := a % b
  return q, r
}


func main()  {
	fmt.Println("Hello World!")
  a, b := div(19, 7) 
  c := func (x, y int) int {
    return x + y + 2
  }
  d := 10
  n := 3
  if d % 2 == 0 {
    fmt.Println("even number!")
  }
  switch n {
    case 1, 2:
      fmt.Println("1 or 2")
    case 3, 4:
      fmt.Println("3 or 4")
  }
  fmt.Println(a, b) // 2, 5
  fmt.Println(c(19, 2))
}