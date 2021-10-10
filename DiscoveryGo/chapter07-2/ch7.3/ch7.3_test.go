package main

import (
	"fmt"
)

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range Chain(PlusOne, PlusOne)(c) {
		fmt.Println(num)
	}
	//Output:
	//.
}

func ExampleFanIn() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() {
		defer close(c1)
		c1 <- 1
		c1 <- 2
	}()
	go func() {
		defer close(c2)
		c2 <- 3
		c2 <- 4
	}()
	go func() {
		defer close(c3)
		c3 <- 5
		c3 <- 6
	}()

	for fi := range FanIn(c1, c2, c3) {
		fmt.Println(fi)
	}
	//Output:
	//.
}
