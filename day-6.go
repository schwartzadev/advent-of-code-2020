package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/juliangruber/go-intersect"
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
	groupSeparator := regexp.MustCompile(`\n\n`)

	groups := groupSeparator.Split(text, -1)

	// Part one.

	// The sum of all of the group counts.
	sumCounts := 0

	for _, group := range groups {
		sumCounts += countUniqueCharacters(group)
	}

	// Print the total count.
	fmt.Println(sumCounts)

	// Part two.

	// The sum of all of the common yes questions per group.
	sumCommon := 0

	// Separates respondents from each other.
	respondentsSeparator := regexp.MustCompile(`\n`)

	for _, group := range groups {
		respondents := respondentsSeparator.Split(group, -1)
		firstResponseChars := stringToCharsArray(respondents[0])

		// This will just be the first response's values.
		commonChars := intersect.Simple(firstResponseChars, firstResponseChars)

		for _, response := range respondents {
			responseAsChars := stringToCharsArray(response)
			commonChars = intersect.Simple(commonChars, responseAsChars)
		}
		// commonChars should now only contain the questions to which all
		// respondents in the group answered yes.
		sumCommon += len(commonChars)
	}

	fmt.Println(sumCommon)
}

func stringToCharsArray(input string) []string {
	return strings.Split(input, "")
}

// Count the number of unique characters in a string. (Excludes newlines.)
func countUniqueCharacters(input string) int {
	characters := stringToCharsArray(input)
	uniqueChars := unique(characters)
	charsNoNewlines := removeNewLines(uniqueChars)
	return len(charsNoNewlines)
}

// Removes the newline character from a slice of strings.
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
