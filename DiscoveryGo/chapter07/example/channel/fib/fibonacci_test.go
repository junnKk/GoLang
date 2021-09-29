package fibonacci

import "fmt"

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ",")
	}
	// Output:
	// 0,1,1,2,3,5,8,13,
}

func ExampleFibonacciGenerator() {
	fib := FibonacciGenerator(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Print(n, ",")
	}
	// Output:
	// 0,1,1,2,3,5,8,13,
}
