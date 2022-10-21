package main

import "fmt"

/*
Hands-on exercise #5
● Show the comma ok idiom starting with this code.

code:
func main() {
	c := make(chan int)

	v, ok :=
	fmt.Println(v, ok)

	close(c)

	v, ok =
	fmt.Println(v, ok)
}

*/

func main() {
	c := make(chan int)

	go func() {
		c <- 5
		c <- 7
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
