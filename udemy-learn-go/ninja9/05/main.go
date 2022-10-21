package main

// Hands-on exercise #5
// Fix the race condition you created in exercise #4 by using package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var count int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			r := atomic.LoadInt64(&count)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", count)
}
