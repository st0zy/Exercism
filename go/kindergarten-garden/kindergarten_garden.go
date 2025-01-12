package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

var vegMap = map[string]string{
	"V": "violets",
	"C": "clover",
	"R": "radishes",
	"G": "grass",
}

type Garden struct {
	children    map[string]int
	plantMatrix [][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	diagram = strings.TrimLeft(diagram, "\n")
	rows := strings.Split(diagram, "\n")
	if len(rows) != 2 {
		return nil, errors.New("Wrong Diagram")
	}
	sort.Strings(children)

	rowLength := len(rows[0])
	if rowLength != len(rows[1]) {
		return nil, errors.New("Mismatched rows")
	}
	if rowLength%2 != 0 {
		return nil, errors.New("Odd number of cups")
	}

	g := &Garden{
		children:    make(map[string]int),
		plantMatrix: make([][]string, rowLength),
	}
	for i, child := range children {
		_, ok := g.children[child]
		if ok {
			return nil, errors.New("Duplicate names detected")
		}
		g.children[child] = i
	}

	// fmt.Println(rowLength)
	for i := 0; i < 2; i++ {
		g.plantMatrix[i] = make([]string, rowLength)

		for j, ch := range rows[i] {
			veg, ok := vegMap[string(ch)]
			if !ok {
				return nil, errors.New("Incorrect cupCode")
			}
			g.plantMatrix[i][j] = veg
		}
	}

	return g, nil

}

func (g *Garden) Plants(child string) ([]string, bool) {
	position, ok := g.children[child]
	if !ok {
		return nil, false
	}

	mapChars := []string{g.plantMatrix[0][2*position],
		g.plantMatrix[0][2*position+1],
		g.plantMatrix[1][2*position],
		g.plantMatrix[1][2*position+1],
	}

	return mapChars, true

}
