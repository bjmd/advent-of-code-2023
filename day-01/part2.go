package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func partTwo(input []string) {

	var total int
	// Regex to parse out both the spelled out numbers and digits
	// Only match from start of string, to prevent overlapping words getting missed
	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for _, line := range input {

		var currentLineNumbers []string

		for i := range line {
			// Drop the first char on each loop. This is to prevent overlapping words getting missed
			found := re.FindString(line[i:])
			// Don't want to append empty strings to our slice
			if found != "" {
				currentLineNumbers = append(currentLineNumbers, found)
			}
		}

		// Converts "one" to "1", etc.
		for index, word := range currentLineNumbers {
			if !isNumber(word) && word != "" {
				currentLineNumbers[index] = wordToNum(word)
			}
		}

		// Grabs and formats 1st and last strings in line to a double digit number
		first := currentLineNumbers[0]
		last := currentLineNumbers[len(currentLineNumbers)-1]

		number, _ := strconv.Atoi(first + last)

		total += number
	}
	fmt.Printf("Part Two Total %d\n", total)
}

// Check if the string is a number by returning whether or not there was
// an error during type conversion
func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

// Converts spelled out numbers to string of digit
// e.g. "one" becomes "1"
func wordToNum(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return ""
	}
}
