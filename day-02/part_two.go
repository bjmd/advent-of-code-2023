package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertStrToIntP2(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func calcHighestOption(revealList []map[string]int) map[string]int {
	highestOptionMap := make(map[string]int)

	for _, reveal := range revealList {
		for colour, count := range reveal {
			if highestOptionMap[colour] < count {
				highestOptionMap[colour] = count
			}
		}
	}

	return highestOptionMap

}

func parseColoursP2(reveal string) map[string]int {
	colourMap := make(map[string]int)

	reColour := regexp.MustCompile(`[a-z]+`)
	reCount := regexp.MustCompile(`\d+`)

	colours := reColour.FindAllString(reveal, -1)
	counts := reCount.FindAllString(reveal, -1)

	for i := 0; i < len(colours); i += 1 {
		count := convertStrToIntP2(counts[i])
		colourMap[colours[i]] += count
	}

	return colourMap
}

func parseRevealsP2(line string) []map[string]int {
	var revealList []map[string]int

	for _, reveal := range strings.Split(line, ";") {
		revealList = append(revealList, parseColoursP2(reveal))
	}

	return revealList

}

func partTwo(input []string) {

	var fewestCubesMap map[string]int
	var fewestCubesPower int
	var fewestCubesSum int
	parsedGameData := make(map[int][]map[string]int)

	for _, line := range input {
		gameSplit := strings.Split(line, ":")
		gameNum := convertStrToInt(strings.Split(gameSplit[0], " ")[1])

		parsedGameData[gameNum] = parseRevealsP2(gameSplit[1])
		fewestCubesMap = calcHighestOption(parsedGameData[gameNum])
		fewestCubesPower = fewestCubesMap["blue"] * fewestCubesMap["green"] * fewestCubesMap["red"]
		fewestCubesSum += fewestCubesPower
	}

	// for game, reveals := range parsedGameData {
	// 	gamePossible := true

	// 	for _, reveal := range reveals {
	// 		for color, count := range reveal {
	// 			if count > coloursToCheck[color] {
	// 				gamePossible = false
	// 				break
	// 			}
	// 		}
	// 	}
	// 	if gamePossible {
	// 		gameRunningCount += game
	// 	}
	// }
	fmt.Printf("Fewest Cubes Sum: %d\n", fewestCubesSum)
}
