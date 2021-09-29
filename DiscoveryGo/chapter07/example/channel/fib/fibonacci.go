package fibonacci

// 피보나치 수열을 max까지 생성합니다.
func Fibonacci(max int) <-chan int { // 채널 이용
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func FibonacciGenerator(max int) func() int { // 클로저 이용
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}
