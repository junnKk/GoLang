package main

import "fmt"

func f() (i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f():", r)
			i = -1
		}
	}()
	g()
	return 100
}

func g() {
	panic("g(): I panic!!")
}

func ExamplePanic() {
	fmt.Println(f())
	// Output:
	// Recovered in f(): g(): I panic!!
	// -1
}
