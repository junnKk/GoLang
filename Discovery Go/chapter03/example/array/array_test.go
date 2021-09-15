package array

import (
	"fmt"

	"github.com/glowingedge/discoveryGo/chapter3/hangul"
)

func Example_array() {
	fruits := [...]string{"apple", "banana", "tomato"}
	for _, fruit := range fruits {
		fmt.Printf("%s is delicious.\n", fruit)
	}

	// Output:
	// apple is delicious.
	// banana is delicious.
	// tomato is delicious.
}

func Example_arrayWithSuffix() {
	fruits := [...]string{"사과", "바나나", "토마토", "멜론"}
	for _, fruit := range fruits {
		suffix := "는"
		if hangul.HasConsonantSuffix(fruit) {
			suffix = "은"
		}
		fmt.Printf("%s%s 맛있어.\n", fruit, suffix)
	}

	// Output:
	// 사과는 맛있어.
	// 바나나는 맛있어.
	// 토마토는 맛있어.
	// 멜론은 맛있어.
}
