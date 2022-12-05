package queenattack

import "errors"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {

	if whitePosition == "" || blackPosition == "" {
		return false, errors.New("empty position")
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}

	lw := string(whitePosition[0])
	lwi := m[lw]
	nw := int(whitePosition[1] - '0')

	lb := string(blackPosition[0])
	lbi := m[lb]
	nb := int(blackPosition[1] - '0')

	switch {
	case len(whitePosition) != 2 || len(blackPosition) != 2:
		return false, errors.New("invalid position")
	case nw < 1 || nw > 8 || nb < 1 || nb > 8 || lw < "a" || lw > "h" || lb < "a" || lb > "h":
		return false, errors.New("position off board")
	case whitePosition == blackPosition:
		return false, errors.New("same square")
	case lw == lb:
		return true, nil
	case nw == nb:
		return true, nil
	case (lwi-nw == lbi-nb) || (lwi+nw == lbi+nb):
		return true, nil
	default:
		return false, nil
	}

}
