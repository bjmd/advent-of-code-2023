package main

import (
	"bufio"
	"log"
	"os"
)

// const inputFile = "test_input.txt"

const inputFile = "input.txt"

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	lines, err := readLines(inputFile)

	if err != nil {
		log.Fatalf("readlines: %s", err)
	}

	partOne(lines)
	partTwo(lines)

}
