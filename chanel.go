package main

import (
	"fmt"
)

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main()  {
	ch := make(chan int)
	go receiver(ch)
	i := 0
	for i < 10 {
		ch <- i
		i++
	}
}

/*
0
1
2
3
4
5
6
7
8
9
*/