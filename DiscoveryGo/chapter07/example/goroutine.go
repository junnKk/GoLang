package main

import "fmt"

func main() {
	go func() {
		fmt.Println("In goroutine!")
	}()
	func() {
		fmt.Println("In main routine!")
	}()
}
