package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func convertStrToInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

func calcWinning(count int) int {

	number := 1

	for i := 0; i < count - 1; i++ {
		number = number * 2
	}

	return number

}

func partOne(input []string) {

	runningTotal := 0
	
	for _, line := range input {
		winningCount := 0
		gameSplit := strings.Split(line, ":")
		chosenNumbers := strings.Fields(strings.Split(gameSplit[1], "|")[0])
		winningNums := strings.Fields(strings.Split(gameSplit[1], "|")[1])
		// gameNum := convertStrToInt(strings.Split(gameSplit[0], " ")[1])

		for _, num := range chosenNumbers {
			if slices.Contains(winningNums, num) {
				winningCount += 1
			}
		}

		if winningCount > 0 {
			runningTotal += calcWinning(winningCount)
		}

		// fmt.Printf("Num: %d\n", gameNum)
		// fmt.Printf("Numbers: %s, Winning Numbers %s\n", chosenNumbers, winningNums)

	}
	fmt.Printf("Winning Count %d\n", runningTotal)
}