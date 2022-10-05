package main

// Hands-on exercise #2
// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

import "fmt"

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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favIcecreams {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
