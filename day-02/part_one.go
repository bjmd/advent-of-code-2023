package main

import (
	"fmt"
	"regexp"
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

func parseColours(reveal string) map[string]int {
	colourMap := make(map[string]int)

	reColour := regexp.MustCompile(`[a-z]+`)
	reCount := regexp.MustCompile(`\d+`)

	colours := reColour.FindAllString(reveal, -1)
	counts := reCount.FindAllString(reveal, -1)

	for i := 0; i < len(colours); i += 1 {
		count := convertStrToInt(counts[i])
		colourMap[colours[i]] += count
	}

	return colourMap
}

func parseReveals(line string) []map[string]int {
	var revealList []map[string]int

	for _, reveal := range strings.Split(line, ";") {
		revealList = append(revealList, parseColours(reveal))
	}

	return revealList

}

func partOne(input []string) {

	var gameRunningCount int
	parsedGameData := make(map[int][]map[string]int)
	coloursToCheck := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, line := range input {
		gameSplit := strings.Split(line, ":")
		gameNum := convertStrToInt(strings.Split(gameSplit[0], " ")[1])

		parsedGameData[gameNum] = parseReveals(gameSplit[1])
	}

	for game, reveals := range parsedGameData {
		gamePossible := true

		for _, reveal := range reveals {
			for color, count := range reveal {
				if count > coloursToCheck[color] {
					gamePossible = false
					break
				}
			}
		}
		if gamePossible {
			gameRunningCount += game
		}
	}
	fmt.Printf("Game count: %d\n", gameRunningCount)
}
