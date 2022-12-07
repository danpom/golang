package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	res := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					res = append(res, Triplet{a, b, c})
				}
			}
		}
	}
	return res
}

//below solution doesn't find all triplets
// for i := min; i <= max; i++ {
// 	if i%2 == 0 {
// 		x := int(math.Pow(float64(i/2), 2) - 1)
// 		y := int(math.Pow(float64(i/2), 2) + 1)

// 		if x > i && y > x && math.Pow(float64(i), 2)+math.Pow(float64(x), 2) == math.Pow(float64(y), 2) && y <= max {
// 			triplets = append(triplets, Triplet{i, x, y})
// 		}

// 	} else {
// 		x := int(math.Pow(float64(i), 2)/2 - 0.5)
// 		y := int(math.Pow(float64(i), 2)/2 + 0.5)

// 		if x > i && y > x && math.Pow(float64(i), 2)+math.Pow(float64(x), 2) == math.Pow(float64(y), 2) && y <= max {
// 			triplets = append(triplets, Triplet{i, x, y})
// 		}
// 	}
// }
// return triplets
//		}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	triplets := Range(1, p/2)
	var output []Triplet

	for _, v := range triplets {
		if v[0]+v[1]+v[2] == p {
			output = append(output, v)
		}
	}

	return output
}
