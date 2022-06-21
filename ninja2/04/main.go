package main

import "fmt"

func main() {
	const (
		a int = 10
	)

	fmt.Printf("%d\t%b\t%x", a, a, a)

	b := a << 1

	fmt.Printf("\n%d\t%b\t%x", b, b, b)

}
