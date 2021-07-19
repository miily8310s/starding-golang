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

	type Feed struct {
    Name string
    Amount uint
  }
  
  type Animal struct {
    Name string
    Feed Feed
  }
  
  a := Animal{
    Name: "Monkey",
    Feed: Feed{
      Name: "Banana",
      Amount: 10,
    },
  } 
  
  fmt.Println(a.Name)
  fmt.Println(a.Feed.Name)

	type Error interface {
    Error() string
  }
  
  func DoSomething() (int, error) {
    return int , error
  } 
  _, err := DoSomething()
  if err != nil {
    fmt. Println(err.Error()) 
  }
}