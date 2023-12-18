package main

import (
	"fmt"
	"slices"
	"strings"
)

func partTwo(input []string) {

	cardCount := make(map[int]int)
	var totalCards int
	
	for _, line := range input {
		winningCount := 0
		gameSplit := strings.Split(line, ":")
		chosenNumbers := strings.Fields(strings.Split(gameSplit[1], "|")[0])
		winningNums := strings.Fields(strings.Split(gameSplit[1], "|")[1])
		gameNum := convertStrToInt(strings.Fields(gameSplit[0])[1])

		cardCount[gameNum] += 1

 		for _, num := range chosenNumbers {
			if slices.Contains(winningNums, num) {
				winningCount += 1
			}
		}

		if winningCount > 0 {
			for x := 0; x < cardCount[gameNum]; x++ {
				for i := gameNum + 1; i <= gameNum + winningCount; i++ {
					cardCount[i] += 1
				}
			}
		}
	}
	for _, numOfCards := range cardCount {
		totalCards = totalCards + numOfCards
	}
	fmt.Printf("Winning Count %d\n", totalCards)
}