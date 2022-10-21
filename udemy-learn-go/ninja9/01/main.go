package main

//Hands-on exercise #1
// in addition to the main goroutine, launch two additional goroutines
// ○ each additional goroutine should print something out
// ● use waitgroups to make sure each goroutine finishes before your program exists

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Before")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("function one talking")
		wg.Done()
	}()
	fmt.Println("During?")
	go func() {
		fmt.Println("function two talking")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("After")
}
