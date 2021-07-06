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
	fmt.Println(a, b) // 2, 5
}