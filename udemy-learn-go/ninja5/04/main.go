package main

// Hands-on exercise #4
// Create and use an anonymous struct

import "fmt"

func main() {

	s := struct {
		brand   string
		model   string
		colours []string
	}{
		brand: "Raybans",
		model: "Aviator",
		colours: []string{
			"Green",
			"Black",
			"Blue",
		},
	}
	fmt.Println(s.brand)
	fmt.Println(s.model)
	fmt.Println(s.colours)

	for i, val := range s.colours {
		fmt.Println(i, val)
	}
}
