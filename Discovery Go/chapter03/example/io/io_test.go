package io

import (
	"fmt"
	"os"
	"strings"
)

func Example_input() {
	f, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("open error")
		// Error handling
	}
	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
		fmt.Println(num)
		// Do somethig with data
	}

	// Output:
	// 3
}

func Example_write() {
	f, err := os.Create("abcd.txt")
	if err != nil {
		fmt.Println("crate error")
		// error handling
	}
	defer f.Close()
	num := 3
	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
		fmt.Println("write error")
		// error handling
	}

	// Output:
}

func ExampleWriteTo() {
	lines := []string{
		"bill@e.com",
		"tom@e.com",
		"jane@e.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}

	// Output:
	// bill@e.com
	// tom@e.com
	// jane@e.com
}

func ExampleReadFrom() {
	r := strings.NewReader("bil\ntom\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bil tom jane]
}
