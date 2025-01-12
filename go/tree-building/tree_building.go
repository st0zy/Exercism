package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	err := validateRecords(records)
	if err != nil {
		return nil, err
	}

	var nodes []*Node

	for i, _ := range records {
		node := &Node{
			records[i].ID,
			make([]*Node, 0),
		}
		nodes = append(nodes, node)
	}

	for i, _ := range records {

		if records[i].Parent != i {
			nodes[records[i].Parent].Children = append(nodes[records[i].Parent].Children, nodes[i])
		}
	}

	for i, _ := range records {
		if records[i].Parent == i {
			return nodes[i], nil
		}
	}

	return nil, fmt.Errorf("Something went wrong.")
}

func validateRecords(records []Record) error {
	rootFound := false

	for i, _ := range records {
		if i != 0 && records[i].Parent >= i {
			return fmt.Errorf("Parent ID greater than child")
		}
		if records[i].ID >= len(records) {
			return fmt.Errorf("missing records")
		}
		if i != records[i].ID {
			return fmt.Errorf("duplicate IDs in input")
		}
		if records[i].Parent == i {
			if rootFound {
				return fmt.Errorf("duplicate root found")
			}
			rootFound = true
		}
	}

	if !rootFound {
		return fmt.Errorf("root does not exist")
	}

	return nil
}
