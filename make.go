package main

import (
	"fmt"
)

func main()  {
	s := make([]int, 10)
	a := s[0:4]
	fir := [] int {1, 2, 3}
	fir = append(fir, 4)
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "China"
	m[47] = "Japan"
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(fir)
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}
}