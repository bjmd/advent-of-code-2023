package main

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/pp/v3"
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

func parseLine(line string) (map[int]map[int][]int, []int) {
	var symbolSlice []int 

	numberMap := make(map[int]map[int][]int)

	var currentNum string
	var numPositionSlice []int
	
	for position, item := range line {
		if isNumber(string(item)) {
			numPositionSlice = append(numPositionSlice, position)
			currentNum += string(item)
		} else {
			if len(currentNum) > 0 {
				numberMap[position] = make(map[int][]int)
				numberMap[position][convertStrToInt(currentNum)] = numPositionSlice
			}
			if string(item) != "." {
				symbolSlice = append(symbolSlice, position)
			}
			currentNum = ""
			numPositionSlice = nil

		}
	}

	return numberMap, symbolSlice

}

func positionTouchesNumber(positionList []int, symbolPosition int) bool {

	touching := false	

	for _, position := range positionList {
		if position == symbolPosition || 
			position == symbolPosition - 1 ||
			position == symbolPosition + 1 {

			touching = true

		}
	}

	return touching
	
}

func checkMap(numberMap map[int]map[int][]int, runningTotal int, symbolPosition int) (int, map[int]map[int][]int) {
	for _, linePosition := range numberMap {
		for number, positionList := range linePosition {
			if positionTouchesNumber(positionList, symbolPosition) {
				delete(linePosition, number)
				runningTotal += number
			}
		}
	}
	return runningTotal, numberMap
}

func partOne(input []string) {

	var runningTotal int
	var lineBuffer []string
	var newBuffer []string
	
	firstRun := true
	debug := false
	lineZeroNumberMap := make(map[int]map[int][]int)
	lineOneNumberMap := make(map[int]map[int][]int)
	lineTwoNumberMap := make(map[int]map[int][]int)

	for _, lineString := range input {

		lineBuffer = append(lineBuffer, lineString)

		if debug {
			fmt.Println("\nBuffer")
			pp.Print(lineBuffer)
			fmt.Println("")

		}

		if len(lineBuffer) == 3 {
			if firstRun {
				lineZeroNumberMap, _ = parseLine(lineBuffer[0])
				lineOneNumberMap, _ = parseLine(lineBuffer[1])
				firstRun = false
			}

			lineTwoNumberMap, _ = parseLine(lineBuffer[2])

			for lineNumber, lineString := range lineBuffer {

				_, currentLineSymbols := parseLine(lineString)
				if debug {pp.Print((currentLineSymbols))}

				if len(currentLineSymbols) > 0 {
					if lineNumber == 0 {
						if debug {fmt.Println("Hit 0")}

						for _, symbolPosition := range currentLineSymbols {
							runningTotal, lineZeroNumberMap = checkMap(lineZeroNumberMap, runningTotal, symbolPosition)
							runningTotal, lineOneNumberMap = checkMap(lineOneNumberMap, runningTotal, symbolPosition)
						}
						if debug {fmt.Printf("Running total: %d\n", runningTotal)}
					}
					
					if lineNumber == 1 {
						if debug {fmt.Println("Hit 1")}
						
						for _, symbolPosition := range currentLineSymbols {
							runningTotal, lineZeroNumberMap = checkMap(lineZeroNumberMap, runningTotal, symbolPosition)
							runningTotal, lineOneNumberMap = checkMap(lineOneNumberMap, runningTotal, symbolPosition)
							runningTotal, lineTwoNumberMap = checkMap(lineTwoNumberMap, runningTotal, symbolPosition)
						}
						if debug {fmt.Printf("Running total: %d\n", runningTotal)}
					}
					
					if lineNumber == 2 {
						if debug {fmt.Println("Hit 2")}
						
						for _, symbolPosition := range currentLineSymbols {
							runningTotal, lineOneNumberMap = checkMap(lineOneNumberMap, runningTotal, symbolPosition)
							runningTotal, lineTwoNumberMap = checkMap(lineTwoNumberMap, runningTotal, symbolPosition)
						}
						if debug {fmt.Printf("Running total: %d\n", runningTotal)}
					}
				}			
			}

			newBuffer = nil
			newBuffer = append(newBuffer, lineBuffer[1])
			newBuffer = append(newBuffer, lineBuffer[2])
	
			lineBuffer = newBuffer

			lineZeroNumberMap = lineOneNumberMap
			lineOneNumberMap = lineTwoNumberMap
		}
	}
	fmt.Printf("Total: %d\n", runningTotal)
}
