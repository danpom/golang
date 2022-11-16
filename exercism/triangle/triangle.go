/*Package triangle is used to determine the kind of triangle (equilateral, isosceles or scalene) given the length of its' 3 sides
 */
package triangle

import "sort"

// indicates the kind of triangle we're dealing with e.g. equilateral, isosceles or scalene
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides() determines the kind of triangle (equilateral, isosceles or scalene) given the length of its' 3 sides
func KindFromSides(a, b, c float64) Kind {
	//side labels are arbritrary. Need to organise by length
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	switch {
	case sides[2] >= sides[0]+sides[1]:
		return NaT
	case sides[0] <= 0:
		return NaT
	case sides[0] == sides[2]:
		return Equ
	case sides[1] == sides[2]:
		return Iso
	case sides[0] != sides[1]:
		return Sca
	default:
		return NaT
	}
}
