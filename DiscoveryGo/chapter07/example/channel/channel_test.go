package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 이 함수는 영원히 끝나지 않음
// func Example_simpleChannel3() {
// 	c := make(chan int)
// 	go func() {
// 		c <- 1
// 		c <- 2
// 	}()
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_simpleChannel2() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Plusone(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}
func ExamplePlusone() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range Plusone(Plusone(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func ExampleChain() {
	PlusTwo := Chain(Plusone, Plusone)
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusTwo(c) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

func ExampleFanOut() {
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
	wg.Add(len(ins))
	for _, in := range ins {
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

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}
	timeout := time.After(5 * time.Second)
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
			case <-timeout:
				return
			}
		}
	}()
	return out
}

func ExampleFanIn3() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)
	for n := range FanIn3(c1, c2, c3) {
		fmt.Println(n, " ")
	}
}

func PlusoneWithDone(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-done:
				return
			}
		}
	}()
	return out
}

func ExamplePlusoneWithDone() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	done := make(chan struct{})
	nums := PlusoneWithDone(done, PlusoneWithDone(done, PlusoneWithDone(done, PlusoneWithDone(done, PlusoneWithDone(done, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	close(done)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	for range nums {
		// consume the remain
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}

func PlusoneWithCtx(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func ExamplePlusoneWithCtx() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	nums := PlusoneWithCtx(ctx, PlusoneWithCtx(ctx, PlusoneWithCtx(ctx, PlusoneWithCtx(ctx, PlusoneWithCtx(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			cancel()
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	for range nums {
		// consume the remain
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}
