package main

import (
	"fmt"
)

func main() {
	fib(5)
}

func fib(n int) {
	p, q := 0, 1
	for i := 0; i < n-2; i++ {
		p, q = q, p+q
	}
	fmt.Println(q)
}
