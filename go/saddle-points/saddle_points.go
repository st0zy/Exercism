package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.

type Pair struct {
	row, column int
}

type Matrix struct {
	data [][]int
	rows int
	cols int
}

func New(s string) (*Matrix, error) {
	m := &Matrix{}

	if s == "" {
		return m, nil
	}
	rows := strings.Split(s, "\n")
	m.data = make([][]int, len(rows))
	m.rows = len(rows)

	for i := 0; i < m.rows; i++ {
		rows[i] = strings.TrimSpace(rows[i])
		vals := strings.Split(rows[i], " ")
		m.cols = len(vals)
		// m.data[i] = make([]int, len(vals))
		for _, val := range vals {
			val, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("incorrect matrix val")
			}
			m.data[i] = append(m.data[i], val)
		}
	}

	if !m.validDimension() {
		return nil, fmt.Errorf("invalid matrix dimension")
	}

	return m, nil

}

func (m *Matrix) validDimension() bool {
	for i := 0; i < m.rows; i++ {
		if len(m.data[i]) != m.cols {
			return false
		}
	}
	return true
}

func (m *Matrix) Saddle() []Pair {

	pair := make([]Pair, 0)

	saddle_row := true
	saddle_col := true

	for r, row := range m.Rows() {

		for c, item := range row {

			for _, v := range m.Rows()[r] {
				if item < v {
					saddle_row = false
				}
			}

			for _, v := range m.Cols()[c] {
				if item > v {
					saddle_col = false
				}
			}

			if saddle_row && saddle_col {
				pair = append(pair, Pair{r + 1, c + 1})
			}
			saddle_row = true
			saddle_col = true

		}

	}

	return pair
}

func (m *Matrix) Rows() [][]int {
	var rows = make([][]int, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < len(m.data[i]); j++ {
			rows[i] = append(rows[i], m.data[i][j])
		}
	}

	return rows
}

func (m *Matrix) Cols() [][]int {
	var cols = make([][]int, m.cols)

	for i := 0; i < m.cols; i++ {
		for j := 0; j < m.rows; j++ {
			cols[i] = append(cols[i], m.data[j][i])
		}
	}

	return cols
}
