package hamming

import "errors"

//calculates the hamming difference between string a and b
func Distance(a, b string) (int, error) {
	alen := len(a)
	blen := len(b)

	// switch {
	// case a == b && b == "":
	// 	return 0, errors.New("empty strands")
	// case :

	// }

	if alen == blen {
		sum := 0
		for i := 0; i < alen; i++ {
			if a[i] != b[i] {
				sum++
			}
		}

		return sum, nil
	} else {
		return 0, errors.New("input strings must have the same number characters")
	}
}
