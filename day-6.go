package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// Day6 contains the answer for Day 6's challenge.
func Day6() {
	path := "day-6-input.txt"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	// Separates groups from each other.
	pattern := regexp.MustCompile(`\n\n`)

	groups := pattern.Split(text, -1)

	// Part one.

	// The sum of all of the group counts.
	sumCounts := 0

	for _, group := range groups {
		sumCounts += countUniqueCharacters(group)
	}

	// Print the total count.
	fmt.Println(sumCounts)
}

func countUniqueCharacters(input string) int {
	characters := strings.Split(input, "")
	uniqueChars := unique(characters)
	charsNoNewlines := removeNewLines(uniqueChars)
	return len(charsNoNewlines)
}

func removeNewLines(characters []string) []string {
	for index, character := range characters {
		if character == "\n" {
			// Remove the newline character if it is present.
			return append(characters[:index], characters[index+1:]...)
		}
	}
	// Otherwise, return the input slice.
	return characters
}

// Returns the values array with only unique elements.
func unique(values []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range values {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
