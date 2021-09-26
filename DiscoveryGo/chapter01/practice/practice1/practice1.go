package main

import (
	"fmt"
)

func main() {
	underwearPrice := 10
	swordPrice := 20
	for i := 0; i < 5; i++ {
		fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", underwearPrice, swordPrice)
		underwearPrice += 10
		swordPrice += 10
	}
}
