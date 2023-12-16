package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func partOne(input []string) {
	var total int

	for _, line := range input {
		firstNumberPattern := regexp.MustCompile(`\d`)
		numbers := firstNumberPattern.FindAllString(line, -1)

		first := numbers[0]
		last := numbers[len(numbers)-1]

		number, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}

		total += number
	}

	fmt.Printf("Part One Total %d\n", total)
}
