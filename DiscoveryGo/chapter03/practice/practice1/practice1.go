/* 어색하지 않은 조사가 붙어서 출력되도록 수정하라 */
package main

import (
	"fmt"

	"github.com/junnKk/golang/discoverygo/chapter03/example/hangul"
)

func Example_array() {
	fruits := [3]string{"사과", "수박", "토마토"}
	for _, fruit := range fruits {
		if hangul.HasConsonantSuffixs(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)

		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)

		}
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func main() {
	Example_array()
}
