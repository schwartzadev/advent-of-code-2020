package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func verifyPassword(password string) bool {
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

// Day2 contains the answer for Day 2's challenge.
func Day2() {
	path := "day-2-input.txt"

	buf, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	snl := bufio.NewScanner(buf)

	count := 0

	// This iterates over each line.
	for snl.Scan() {
		if verifyPassword(snl.Text()) {
			count++
		}
		// fmt.Println(snl.Text())
	}

	fmt.Println(count)

	err = snl.Err()

	if err != nil {
		log.Fatal(err)
	}
}
