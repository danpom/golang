package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Hands-on exercise #3
// ●
// Using goroutines, create an incrementer program
// ○ have a variable to hold the incrementer value
// ○ launch a bunch of goroutines
// ■ each goroutine should
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 65●
// read the incrementer value
// ○ store it in a new variable
// ● yield the processor with runtime.Gosched()
// ● increment the new variable
// ● write the value in the new variable back to the incrementer
// variable
// ● use waitgroups to wait for all of your goroutines to finish
// ● the above will create a race condition.
// ● Prove that it is a race condition by using the -race flag
// ● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

func main() {
	var wg sync.WaitGroup

	count := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			routineCounter := count
			runtime.Gosched()
			routineCounter++
			count = routineCounter
			fmt.Println("Count during:", count)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Count final:", count)

}
