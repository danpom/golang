package main

import "fmt"

// Hands-on exercise #1
// ●
// get this code working:(https://play.golang.org/p/j-EA6003P0)
// ○ using func literal, aka, anonymous self-executing func
// ○ a buffered channel
//code:
// func main() {
// 	c := make(chan int)

// 	c <- 42

// 	fmt.Println(<-c)
// }

//using func literal, aka, anonymous self-executing func
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
