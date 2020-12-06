package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

// Day4 contains the answer for Day 4's challenge.
func Day4() {
	path := "day-4-input.txt"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	// pattern is a pattern for blank lines.
	pattern := regexp.MustCompile(`\n\n`)

	passports := pattern.Split(text, -1)

	validPassportsCount := 0
	for _, passport := range passports {
		if validatePassport(passport) {
			validPassportsCount++
		}
	}
	fmt.Println(validPassportsCount)
}

func validatePassport(passport string) bool {
	byrPattern := regexp.MustCompile(`byr:(19[2-9][0-9]|200[0-2])(\s|$)`)
	iyrPattern := regexp.MustCompile(`iyr:20(1[0-9]|20)(\s|$)`)
	eyrPattern := regexp.MustCompile(`eyr:20(2[0-9]|30)(\s|$)`)
	hgtPattern := regexp.MustCompile(`hgt:(((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in))(\s|$)`)
	hclPattern := regexp.MustCompile(`hcl:#[0-9a-f]{6}(\s|$)`)
	eclPattern := regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)(\s|$)`)
	pidPattern := regexp.MustCompile(`pid:([0-9]{9})(\s|$)`)

	if !byrPattern.MatchString(passport) {
		return false
	}
	if !iyrPattern.MatchString(passport) {
		return false
	}
	if !eyrPattern.MatchString(passport) {
		return false
	}
	if !hgtPattern.MatchString(passport) {
		return false
	}
	if !hclPattern.MatchString(passport) {
		return false
	}
	if !eclPattern.MatchString(passport) {
		return false
	}
	if !pidPattern.MatchString(passport) {
		return false
	}

	return true
}
