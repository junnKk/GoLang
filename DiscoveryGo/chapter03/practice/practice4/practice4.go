/* 슬라이스를 이용하여 queue 구현*/

package main

import "fmt"

func SliceQueue() {
	var queue []string

	// Enqueue
	queue = append(queue, "Hello ")
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Print(queue[0])

		// Dequeue
		queue[0] = ""
		queue = queue[1:]
	}

}

func main() {
	SliceQueue()
}

// 문제점 -> 용량 관리가 안될 수 있음!!?!??
