package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	//calculate Earth years
	ey := seconds / 31557600
	//return planet years based on input
	switch planet {
	case "Mercury":
		return ey / 0.2408467
	case "Venus":
		return ey / 0.61519726
	case "Earth":
		return ey
	case "Mars":
		return ey / 1.8808158
	case "Jupiter":
		return ey / 11.862615
	case "Saturn":
		return ey / 29.447498
	case "Uranus":
		return ey / 84.016846
	case "Neptune":
		return ey / 164.79132
	default:
		return -1
	}
}
