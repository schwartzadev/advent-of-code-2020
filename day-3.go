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

	var lines [323]string
	// Build lines array from input file here.
	i := 0
	for snl.Scan() {
		// Iterate over each line.
		lines[i] = snl.Text()
		i++
	}

	// Part one answer.
	fmt.Println(getTreesCount(lines, 3, 1))

	// Part two answer.
	fmt.Println(
		getTreesCount(lines, 3, 1) * getTreesCount(lines, 1, 1) * getTreesCount(lines, 5, 1) * getTreesCount(lines, 7, 1) * getTreesCount(lines, 1, 2),
	)

	err = snl.Err()

	if err != nil {
		log.Fatal(err)
	}
}

// getTreesCount counts the number of trees in a particular route
func getTreesCount(rows [323]string, dx int, dy int) int {
	treesCount := 0
	x := 0 // Initial x position.
	y := 0 // Initial y position.

	rowWidth := 30 // Initial row width.

	for true {
		if x > rowWidth {
			// We've hit the end of the row, so we duplicate the row content.
			rows = doubleRowsWidth(rows)
			rowWidth = rowWidth * 2
		}
		if y >= 323 {
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
func doubleRowsWidth(rows [323]string) [323]string {
	for i := 0; i < 323; i++ {
		rows[i] = rows[i] + rows[i] // String concatenation.
	}
	return rows
}
