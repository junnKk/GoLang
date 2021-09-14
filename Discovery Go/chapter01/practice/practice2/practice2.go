package main

import (
	"fmt"
)

func main() {
	Hanoi(3)
}

func Hanoi(n int) {
	fmt.Println("Number of disk :", n)
	Move(n, 1, 2, 3)
}

func Move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}
