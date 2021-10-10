package main

import (
	"context"
	"fmt"
)

type IntPipe func(context.Context, <-chan int) <-chan int

// Range는 정수 생성기로 start부터 시작해서 step만큼 더하면서 무한정 생성한다.
// (start, start + step, start + 2*step, ...)
func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}


// n배수를 걸러내는 파이프라인을 반환
// 클로저를 이용하여 함수를 반환하느느 이유는 파이프라인 함수령을 맞춰서 다른 파이프라인에 연결해서 쓸 수 있게 하기 위함
func FilterMultiple(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
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

func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		fmt.Println(prime, " ")
	}
	fmt.Println()
}


func main() {

}
