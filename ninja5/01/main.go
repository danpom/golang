package main

// Hands-on exercise #1
// Create your own type “person” which
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person which stores the favorite flavors

import (
	"fmt"
)

type person struct {
	first        string
	last         string
	favIcecreams []string
}

func main() {
	p1 := person{
		first: "Hannibal",
		last:  "Barca",
		favIcecreams: []string{
			"vanilla",
			"chocolate",
			"raspberry",
		},
	}

	p2 := person{
		first: "John",
		last:  "Malkovich",
		favIcecreams: []string{
			"rum raison",
			"fish",
			"milk",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIcecreams {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIcecreams {
		fmt.Println(i, v)
	}

}
