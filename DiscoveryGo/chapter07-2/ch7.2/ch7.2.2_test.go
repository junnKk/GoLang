package main

import (
	"fmt"
	// "reflect"
)

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, " ")
	}
	//OutPut:
	//.
}

func ExampleFibonacciGenerator() {
	fib := FibonacciGenerator(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Print(n, " ")
	}
	//OutPut:
	//.
}

func ExampleBabyNames() {
	for n := range BabyNames("준지수형경", "희기이김경") {
		fmt.Println(n)
	}
	//OutPut:
	//.
}
func ExampleBabyNamesGenerator() {
	bn := BabyNamesGenerator("준지수형경", "희기이김경")

	for name := bn(); name != ""; name = bn() {
		fmt.Println(name)
	}

	//Output:
	//.
}
