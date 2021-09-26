package generator

import "fmt"

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen(), gen2())
	fmt.Println(gen(), gen2(), gen())

	// Output:
	// 1 1
	// 2 2 3
}
