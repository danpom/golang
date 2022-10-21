package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Hands-on exercise #7
●
write a program that
○ launches 10 goroutines
■ each goroutine adds 10 numbers to a channel
○ pull the numbers off the channel and print them
*/

//note that this solution launches an additional 2 go routines in the set up so I guess not technically correct but I beleive answers the spirit of the actual question
func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for i := 0; i < 10; i++ {
					c <- i*10 + m
					fmt.Println("ROUTINES", runtime.NumGoroutine())
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println("pulling off value from channel: ", v)
	}
}
