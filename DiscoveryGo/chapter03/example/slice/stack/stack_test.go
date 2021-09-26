package stack

import "fmt"

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("5 + 3"))
	fmt.Println(Eval("5 + 3 - 2"))
	fmt.Println(Eval("3 * ( 5 + 3 * 3 ) / 2"))
	fmt.Println(Eval("3 * ( ( 5 + 3 ) * 3 ) / 2"))
	// Output:
	// 5
	// 8
	// 6
	// 21
	// 36
}
