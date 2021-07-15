package main

import (
	"fmt"
)

func main()  {
	// Pointer
	var i int
	p := & i
	i = 5
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
	*p = 10
	fmt.Println(i)
	pp := & p
	fmt.Printf("%T\n", pp)
	pArr := &[3]int{1, 2, 6} 
  fmt.Println((*pArr)[2])

	type Point struct {
		X, Y int
	}
	var pt Point
	pt.X = 10
	pt.Y = 20
	fmt.Println(pt.X)
	fmt.Println(pt.Y)
}