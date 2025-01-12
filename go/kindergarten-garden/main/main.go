package main

import (
	"fmt"
	"kindergarten"
)

func main() {

	// 	diagram := `
	// VCRRGVRG
	// RVGCCGCV`
	g1, _ := kindergarten.NewGarden("RC\nGG", []string{"Alice", "Bob", "Charlie", "Dan"})

	fmt.Println(g1.Plants("Alice"))
}
