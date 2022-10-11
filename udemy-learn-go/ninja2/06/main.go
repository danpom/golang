package main

import "fmt"

func main() {
	const (
		a = 2022 + iota
		b = 2022 + iota
		c = 2022 + iota
		d = 2022 + iota
		e = 2022 + iota
	)
	fmt.Println(a, b, c, d, e)
}
