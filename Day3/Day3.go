package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	partOne(input)
	partTwo(input)
}

func readInput() [][]string {
	// Open our input file
	file, err := os.Open("Day3/Day3.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var grid [][]string

	scanner := bufio.NewScanner(file)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Split on nothing to get an array of individual characters from the whole line
		line := strings.Split(text, "")

		// Add the newly split line to our overall grid
		grid = append(grid, line)
	}

	return grid
}

func partOne(input [][]string) {
	slope := Slope {
		3, 1,
	}

	// Establish our current position
	currentY := 0
	currentX := 0

	treesEncountered := 0

	for currentY < len(input) {
		// If our current position is a tree, count it
		if input[currentY][currentX] == "#" {
			treesEncountered++
		}

		// Since we know the pattern repeats indefinitely to the right, we can just wrap our X value if we reach the end
		if currentX + slope.moveRight >= len(input[currentY]) {
			currentX = (currentX + slope.moveRight) % len(input[currentY])
		} else {
			currentX += slope.moveRight
		}

		currentY += slope.moveDown
	}

	fmt.Println(treesEncountered)
}

func partTwo(input [][]string) {
	slopes := []Slope {
		{
			1, 1,
		},
		{
			3,1,
		},
		{
			5,1,
		},
		{
			7,1,
		},
		{
			1,2,
		},
	}

	totalScore := 1

	for _, slope := range slopes {
		currentY := 0
		currentX := 0

		treesEncountered := 0

		for currentY < len(input) {
			if input[currentY][currentX] == "#" {
				treesEncountered++
			}

			if currentX + slope.moveRight >= len(input[currentY]) {
				currentX = (currentX + slope.moveRight) % len(input[currentY])
			} else {
				currentX += slope.moveRight
			}

			currentY += slope.moveDown
		}

		totalScore *= treesEncountered
	}


	fmt.Println(totalScore)
}

type Slope struct {
	moveRight int
	moveDown  int
}
