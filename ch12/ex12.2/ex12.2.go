package main

import "fmt"

const Y int = 3

func main() {
	// x := 5
	// a := [x]int{1, 2, 3, 4, 5} // 변수로 선언된 x는 배열 길이로 사용할 수 없으므로 오류

	b := [Y]int{1, 2, 3}

	// fmt.Print(a)
	fmt.Print(b)

}
