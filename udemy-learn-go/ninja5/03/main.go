package main

import "fmt"

// Hands-on exercise #3
// ● Create a new type: vehicle.
// ○ The underlying type is a struct.
// ○ The fields:
// ■ doors
// ■ color
// ● Create two new types: truck & sedan.
// ○ The underlying type of each of these new types is a struct.
// ○ Embed the “vehicle” type in both truck & sedan.
// ○ Give truck the field “fourWheel” which will be set to bool.
// ○ Give sedan the field “luxury” which will be set to bool. solution
// ● Using the vehicle, truck, and sedan structs:
// ○ using a composite literal, create a value of type truck and assign values to the
// fields;
// ○ using a composite literal, create a value of type sedan and assign values to the
// fields.
// ● Print out each of these values.
// ● Print out a single field from each of these values

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	cyberTruck := truck{
		vehicle: vehicle{
			doors:  5,
			colour: "silver",
		},
		fourwheel: true,
	}

	model3 := sedan{
		vehicle: vehicle{
			doors:  5,
			colour: "red",
		},
		luxury: true,
	}

	fmt.Println(cyberTruck)
	fmt.Println(model3)
	fmt.Println(cyberTruck.colour)
	fmt.Println(model3.luxury)

}
