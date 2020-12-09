package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	partOne(input)
	partTwo(input)
}

func readInput() (records []int) {
	// Open our input file
	file, err := os.Open("Day9/Day9.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records = make([]int, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		convertedValue, _ := strconv.Atoi(text)

		records = append(records, convertedValue)
	}

	return records
}

func partOne(input []int) {
	// Set our preamble and look back values
	preamble := 25
	lookBack := 25

	for i := preamble; i < len(input); i++ {
		// Set the current place in our input
		currentValue := input[i]

		// Slice the previous 'x' values based on look back
		currentSlice := input[i - lookBack:i]

		matchingValues := false

		finish:

		// Double loop through the current slice to try and find two values that sum to the current value
		for _, first := range currentSlice {
			for _, second := range currentSlice {
				// If these two values sum to our current value then this number is fine and we can move on
				if first + second == currentValue {
					matchingValues = true
					break finish
				}
			}
		}

		// If we iterate through all possible sums and can't find a match then this is our failure number
		if !matchingValues {
			fmt.Println(currentValue)
			break
		}
	}
}

func partTwo(input []int) {
	// Set the magic number from the result of part 1
	expectedValue := 375054920

	for i, firstValue := range input {
		// Set our current sum to the first value since we would add that anyways
		currentSum := firstValue

		// Use the first value as starting points for our smallest and largest values
		// This is because they will either be smaller, larger, or already equal at the end
		smallest := firstValue
		largest := firstValue

		// Second counter for adding up the rest
		j := i + 1

		// Loop through the next numbers until our sum is >= to our expected value
		for ;currentSum < expectedValue; j++ {
			secondValue := input[j]

			// Check for smallest and largest
			if secondValue < smallest {
				smallest = secondValue
			}
			if secondValue > largest {
				largest = secondValue
			}

			// Add to our sum
			currentSum += secondValue
		}

		// If our current sum equals our expected value then we've found our answer and can stop running
		if currentSum == expectedValue {
			fmt.Println(smallest + largest)
			break
		} else if currentSum > expectedValue {
			continue
		}
	}
}
