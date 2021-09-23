package main

import (
	"fmt"
	"time"
)

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

func main() {
	fmt.Println("Ladies and gnetlemen!")
	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("I am so excited!")
	})
	timer.Stop()
	CountDown(5)
}
