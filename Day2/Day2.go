package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputLines := readInput()

	partOne(inputLines)
	partTwo(inputLines)
}

func readInput() []InputLine {
	// Open our input file
	file, err := os.Open("Day2/Day2.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var lines []InputLine

	scanner := bufio.NewScanner(file)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		var inputLine InputLine

		// Get each individual line
		// Each line comes in this format: {min}-{max} {character}: {password}
		text := scanner.Text()

		// Split on space to get each individual piece
		splitText := strings.Split(text, " ")

		// Split on a hyphen and then assign the min and max values (format of this piece: x-y)
		minAndMax := strings.Split(splitText[0], "-")

		inputLine.MinimumAppearance, _ = strconv.Atoi(minAndMax[0])
		inputLine.MaximumAppearance, _ = strconv.Atoi(minAndMax[1])

		// Remove the colon from the middle section ({character}:) so we just get the expected character
		inputLine.ExpectedCharacter = strings.Replace(splitText[1], ":", "", -1)

		// The password is the last piece of the line, and doesn't need any additional parsing
		inputLine.ProvidedPassword = splitText[2]

		// Add the newly converted int to our return value
		lines = append(lines, inputLine)
	}

	return lines
}

func partOne(inputLines []InputLine) {
	validCount := 0

	for _, line := range inputLines {
		characterCount := strings.Count(line.ProvidedPassword, line.ExpectedCharacter)

		if characterCount >= line.MinimumAppearance && characterCount <= line.MaximumAppearance {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func partTwo(inputLines []InputLine) {
	validCount := 0

	for _, line := range inputLines {
		// Check the characters at the positions of min and max appearance
		// for part two these aren't actually ranges, but positions in the password that need to be checked
		// We subtract one because the puzzle indexes at 1 instead of 0
		firstIndexValue := string(line.ProvidedPassword[line.MinimumAppearance - 1])
		secondIndexValue := string(line.ProvidedPassword[line.MaximumAppearance - 1])

		// See if the two provided indices match the expected character
		firstIndexAppears := firstIndexValue == line.ExpectedCharacter
		secondIndexAppears := secondIndexValue == line.ExpectedCharacter

		// This is an implementation of XOR in Go because it doesn't have one built in
		// If one index exists but not both then it passes
		// If both or neither index exist then it fails
		if firstIndexAppears != secondIndexAppears {
			validCount++
		}
	}

	fmt.Println(validCount)
}

type InputLine struct {
	MinimumAppearance int
	MaximumAppearance int
	ExpectedCharacter string
	ProvidedPassword  string
}
