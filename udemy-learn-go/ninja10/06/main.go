package main

import "fmt"

/*
Hands-on exercise #6
●
write a program that
○ puts 100 numbers to a channel
○ pull the numbers off the channel and print them
*/

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
			fmt.Println("putting on:", i)
		}
		close(c)
	}()

	for v := range c {
		fmt.Println("pulling off:", v)

	}
}
