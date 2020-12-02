package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputArray := readInput()

	partOne(inputArray)
	partTwo(inputArray)
}

func readInput() []int {
	// Open our input file
	file, err := os.Open("Day1/Day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Convert our string line into an int
		convertedText, _ := strconv.Atoi(text)

		// Add the newly converted int to our return value
		lines = append(lines, convertedText)
	}

	return lines
}

func partOne(inputArray []int) {
	// Double loop over our input and add the values together
	for _, first := range inputArray {
		for _, second := range inputArray {
			// Once we find the two values that sum to 2020 we can return their product
			if first + second == 2020 {
				fmt.Println(first * second)

				return
			}
		}
	}
}

func partTwo(inputArray []int) {
	// Triple loop over our input and add the values together
	for _, first := range inputArray {
		for _, second := range inputArray {
			for _, third := range inputArray {
				// Once we find the two values that sum to 2020 we can return their product
				if first + second + third == 2020 {
					fmt.Println(first * second * third)

					return
				}
			}
		}
	}
}
