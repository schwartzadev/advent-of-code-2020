package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Day5 contains the answer for Day 5's challenge.
func Day5() {
	path := "day-5-input.txt"

	buf, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	snl := bufio.NewScanner(buf)

	// Build seats array from input file.
	var seats [876]string
	i := 0
	for snl.Scan() {
		// Iterate over each line.
		seats[i] = snl.Text()
		i++
	}

	// ids holds the id values for each seat.
	var ids [876]int

	for i := 0; i < len(seats); i++ {
		row := getSeatRow(seats[i])
		column := getSeatColumn(seats[i])
		ids[i] = getSeatID(row, column)
	}

	// Get the max value for part one.
	max := ids[0]
	for _, value := range ids {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)

	// Find the missing value for part two.
	sort.Ints(ids[:]) // Sort the ids.

	for i := 1; i < len(ids)-1; i++ {
		// Check if these values are sequential.
		if ids[i-1]+1 != ids[i] {
			// Add one since the actual value is missing.
			fmt.Println(ids[i-1] + 1)
			return
		}
	}
}

func getSeatRow(seat string) int {
	var minRow float64 = 0
	var maxRow float64 = 127
	for i := 0; i < 7; i++ {
		currentChar := seat[i]
		spread := maxRow - minRow + 1
		if currentChar == 70 {
			// F. Keep the lower half.
			maxRow -= math.Floor(spread / 2)
		}
		if currentChar == 66 {
			// B. Keep the upper half.
			minRow += math.Ceil(spread / 2)
		}
	}
	// Could return either minRow or maxRow since they're equal.
	return int(minRow)
}

func getSeatColumn(seat string) int {
	var minRow float64 = 0
	var maxRow float64 = 7
	for i := 7; i < 10; i++ {
		currentChar := seat[i]
		spread := maxRow - minRow + 1
		if currentChar == 76 {
			// L. Keep the lower half.
			maxRow -= math.Floor(spread / 2)
		}
		if currentChar == 82 {
			// R. Keep the upper half.
			minRow += math.Ceil(spread / 2)
		}
	}
	// Could return either minRow or maxRow since they're equal.
	return int(minRow)
}

func getSeatID(seatRow int, seatColumn int) int {
	return seatRow*8 + seatColumn
}
