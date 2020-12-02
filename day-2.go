package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func verifyPasswordPartOne(password string) bool {
	parts := strings.Fields(password)
	countRange := strings.Split(parts[0], "-")

	// minimum frequency
	min, _ := strconv.Atoi(countRange[0])

	// maximum frequency
	max, _ := strconv.Atoi(countRange[1])

	// the letter whose frequency we're counting
	letterToCount := parts[1][0:1]

	// the password string we're verifying
	passwordString := parts[2]

	// the number of times the letter occurs in the password string
	letterFrequency := strings.Count(passwordString, letterToCount)

	return letterFrequency <= max && letterFrequency >= min
}

func verifyPasswordPartTwo(password string) bool {
	parts := strings.Fields(password)
	countRange := strings.Split(parts[0], "-")

	// first index
	first, _ := strconv.Atoi(countRange[0])

	// second index
	second, _ := strconv.Atoi(countRange[1])

	// the letter whose frequency we're counting
	letterToCount := parts[1][0:1]

	// the string to verify
	passwordString := parts[2]

	// true if one of these matches
	matchesFirst := passwordString[first-1:first] == letterToCount
	matchesSecond := passwordString[second-1:second] == letterToCount

	// != works as XOR
	return matchesFirst != matchesSecond
}

// Day2 contains the answer for Day 2's challenge.
func Day2() {
	path := "day-2-input.txt"

	buf, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	snl := bufio.NewScanner(buf)

	partOneCount := 0
	partTwoCount := 0

	// Part one here:

	// This iterates over each line.
	for snl.Scan() {
		if verifyPasswordPartOne(snl.Text()) {
			partOneCount++
		}
		if verifyPasswordPartTwo(snl.Text()) {
			partTwoCount++
		}
	}

	fmt.Println(partOneCount)
	fmt.Println(partTwoCount)

	err = snl.Err()

	if err != nil {
		log.Fatal(err)
	}
}
