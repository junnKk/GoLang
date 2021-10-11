package main

import (
	"context"
	"fmt"
)

type IntPipe func(context.Context, <-chan int) <-chan int

func Print(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for s := range Sums(ctx) {
		if s > max {
			break
		}
		fmt.Println(s)
	}
}

func Sums(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 1)
		for {
			select {
			case i := <-c:
				c = Filter(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func SplitSum(n int) int {
	result := 0
	for n != 0 {
		result += n % 10
		// fmt.Println(result, "ff")
		n /= 10
		// fmt.Println(n)

	}
	return result
}

func Filter(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := range in {
				// fmt.Println(SplitSum(n), n)
				// fmt.Println(SplitSum(i), i)
				if SplitSum(n) == SplitSum(i) {
					continue
				}
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

func Range(ctx context.Context, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func main() {
	Print(1000)

	// fmt.Println(SplitSum(3))

}
