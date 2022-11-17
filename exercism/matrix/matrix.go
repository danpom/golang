package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	//breaking up into rows
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows))
	//row length check
	rl := -1

	//breaking rows up in to digits then converting each to an int
	for i, r := range rows {
		row := strings.Split(strings.TrimSpace(r), " ")
		if rl == -1 {
			rl = len(row)
		} else if rl != len(row) {
			return nil, errors.New("rows must be of equal length")
		}

		cols := make([]int, len(row))
		for j, c := range row {
			v, err := strconv.Atoi(c)
			if err != nil {
				return nil, errors.New("invalid input")
			}
			cols[j] = v
		}
		m[i] = cols
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	res := make([][]int, len(m[0]))
	if len(m) == 0 {
		return res
	}

	for i, r := range m {
		for j, cell := range r {
			if i == 0 {
				res[j] = make([]int, len(m))
			}
			res[j][i] = cell
		}
	}
	return res
}

func (m Matrix) Rows() [][]int {
	res := make([][]int, len(m))
	for i := range m {
		res[i] = make([]int, len(m[i]))
		copy(res[i], m[i])
	}

	return res
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) {
		return false
	}
	if col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
