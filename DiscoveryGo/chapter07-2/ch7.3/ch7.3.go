package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

type IntPipe func(<-chan int) <-chan int

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func FanOut() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				// time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range ins {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 팬아웃해서 파이프라인을 통과시키고 다시 팬인시키면 분산처리가 된다.
// IntPipe 형태의 함수를 받은 뒤에 n개로 분산처리하는 함수로 돌려주는 함수. 팬아웃과 팬인을 모두 수행.
func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

// Distribute과 Chain을 함께 사용하면 다양한 파이프라인을 구성할 수 있다.
// out := Chain(Cut, Distribute(Chain(Draw, Paint, Decorate),10),Box)(in)

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()
	return out
}

func main() {
	FanOut()
	// c1 := make(chan int)
	// c2 := make(chan int)
	// c3 := make(chan int)
	// sendInts := func(c chan<- int, begin, end int) {
	// 	defer close(c)
	// 	for i := begin; i < end; i++ {
	// 		c <- i
	// 	}
	// }
	// go sendInts(c1, 11, 14)
	// go sendInts(c2, 21, 23)
	// go sendInts(c3, 31, 35)
	// for n := range FanIn3(c1, c2, c3) {
	// 	fmt.Println(n)
	// }
}
