package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

func parseLineP2(line string) (map[int]map[int][]int, []int) {
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

func positionTouchesNumberP2(positionList []int, symbolPosition int) bool {

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

func checkMapP2(numberMap map[int]map[int][]int, touchingNumbers []int, symbolPosition int) ([]int, map[int]map[int][]int) {
	for _, linePosition := range numberMap {
		for number, positionList := range linePosition {
			if positionTouchesNumberP2(positionList, symbolPosition) {
				delete(linePosition, number)
				touchingNumbers = append(touchingNumbers, number)
			}
		}
	}
	return touchingNumbers, numberMap
}

func partTwo(input []string) {

	var runningTotal int
	
	var lineBuffer []string
	var newBuffer []string
	// var allSymbols [][]int
	
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
				lineZeroNumberMap, _ = parseLineP2(lineBuffer[0])
				lineOneNumberMap, _ = parseLineP2(lineBuffer[1])
				firstRun = false
			}
			
			lineTwoNumberMap, _ = parseLineP2(lineBuffer[2])
			
			for lineNumber, lineString := range lineBuffer {
				
				_, currentLineSymbols := parseLineP2(lineString)
				if debug {pp.Print((currentLineSymbols))}
				
				if len(currentLineSymbols) > 0 {
					for _, symbolPosition := range currentLineSymbols {
						if lineNumber == 1 {
							var touchingNumbers []int
							if debug {fmt.Println("Hit 1")}
							touchingNumbers, lineZeroNumberMap = checkMapP2(lineZeroNumberMap, touchingNumbers, symbolPosition)
							touchingNumbers, lineOneNumberMap = checkMapP2(lineOneNumberMap, touchingNumbers, symbolPosition)
							touchingNumbers, lineTwoNumberMap = checkMapP2(lineTwoNumberMap, touchingNumbers, symbolPosition)
							if len(touchingNumbers) == 2 {
								runningTotal += touchingNumbers[0] * touchingNumbers[1]
							}
						}
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
