package main

import (
	"fmt"
)

func main()  {
	s := make([]int, 10)
	a := s[0:4]
	fir := [] int {1, 2, 3}
	fir = append(fir, 4)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(fir)
}