package main

import (
	"fmt"
	"sync"
)

// run with --race option
func main() {
	cnt := int64(10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// wg.Wait()
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt--
		}()
	}
	wg.Wait()
	fmt.Println("cnt:", cnt)
}
