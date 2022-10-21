package main

import "fmt"

// Hands-on exercise #2
// ●
// Get this code running:
// ○https://play.golang.org/p/oB-p3KMiH6

//code:
// func main() {
// 	cs := make(chan<- int)

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
