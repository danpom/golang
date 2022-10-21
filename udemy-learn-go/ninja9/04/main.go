package main

// Hands-on exercise #4
// Fix the race condition you created in the previous exercise by using a mutex
// ● it makes sense to remove runtime.Gosched()

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	count := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			routineCounter := count
			//runtime.Gosched()
			routineCounter++
			count = routineCounter
			fmt.Println("Count during:", count)
			fmt.Println("number of goroutines: ", runtime.NumGoroutine())
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Count final:", count)

}
