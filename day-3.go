package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Day3 contains the answer for Day 3's challenge.
func Day3() {
	path := "day-3-input.txt"

	buf, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	snl := bufio.NewScanner(buf)

	// Build lines array from input file.
	var lines [323]string
	i := 0
	for snl.Scan() {
		// Iterate over each line.
		lines[i] = snl.Text()
		i++
	}

	// Part one answer (right three, down one).
	fmt.Println(getTreesCount(lines[:], 3, 1))

	// Part two answer.
	fmt.Println(
		getTreesCount(lines[:], 3, 1) * getTreesCount(lines[:], 1, 1) * getTreesCount(lines[:], 5, 1) * getTreesCount(lines[:], 7, 1) * getTreesCount(lines[:], 1, 2),
	)

	err = snl.Err()

	if err != nil {
		log.Fatal(err)
	}
}

// getTreesCount counts the number of trees in a particular route
func getTreesCount(rows []string, dx int, dy int) int {
	treesCount := 0
	x := 0 // Initial x position.
	y := 0 // Initial y position.

	rowWidth := len(rows[0]) // Initial row width.

	for true {
		if x > rowWidth {
			// We've hit the end of the row, so we duplicate the row content.
			doubleRowsWidth(rows[:]) // Passes rows as a slice (by reference)
			rowWidth = rowWidth * 2
		}
		if y >= len(rows) {
			return treesCount
		}
		if rows[y][x:x+1] == "#" {
			// This location is a tree.
			treesCount++
		}
		x += dx // Update x.
		y += dy // Update y.
	}
	return -1
}

// doubleRowsWidth appends the content of each row to itself.
func doubleRowsWidth(rows []string) {
	for i := 0; i < len(rows); i++ {
		rows[i] = rows[i] + rows[i] // String concatenation.
	}
}
