package main

import "fmt"

func Example_simpleChannel() {
	ch := func() <-chan int {

		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range ch {
		fmt.Println(num)

	}

}

func main() {
	Example_simpleChannel()
}
