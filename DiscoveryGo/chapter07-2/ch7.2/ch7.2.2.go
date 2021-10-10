package main

import (
	"fmt"
	"strings"
)

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next

	}

}
func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func BabyNamesGenerator(first, second string) func() string {
	name := ""
	firstIndex, secondIndex := 0, 0
	firstA := strings.Split(first, "")
	secondA := strings.Split(second, "")
	return func() string {
		if secondIndex == len(secondA) && firstIndex == len(firstA)-1 {
			name = ""
			return name
		}
		if secondIndex == len(secondA) {
			firstIndex++
			secondIndex = 0
		}
		name = string(firstA[firstIndex]) + string(secondA[secondIndex])
		secondIndex++
		return name
	}

}

func main() {
	// for fib := range Fibonacci(15) {
	// 	fmt.Print(fib, " ")
	// }

	// fib := FibonacciGenerator(15)
	// for n := fib(); n >= 0; n = fib() {
	// 	fmt.Print(n, " ")
	// }
	bn := BabyNamesGenerator("준지수형경", "희기이김경")
	
	fmt.Println(bn())
		
}
