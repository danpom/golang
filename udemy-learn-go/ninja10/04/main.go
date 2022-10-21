package main

import "fmt"

/*
Hands-on exercise #4
● Starting with this code, pull the values off the channel using a select statement

code:
func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

*/

func main() {
	c, q := gen()
	receive(c, q)

	fmt.Println("about to exit")
}

func gen() (<-chan int, <-chan int) {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(q)
	}()

	return c, q
}

func receive(c, q <-chan int) {
	var x int
	for {
		select {
		case x = <-c:
			fmt.Println(x)
		case <-q:
			return
		}
	}
}
