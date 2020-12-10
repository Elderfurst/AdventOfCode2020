package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readInput()

	partOne(input)
	partTwo(input)
}

func readInput() (records []int) {
	// Open our input file
	file, err := os.Open("Day10/Day10.txt")

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

	// Sort our input so it's easier to work with for the challenge
	sort.Ints(records)

	return records
}

func partOne(input []int) {
	// The 'charger' starts at 0 jolts, so we start there too
	currentJolts := 0

	differences := make (map[int]int, 0)

	// Since our device is always 3 higher than the highest value
	// We default to a minimum of one '3' difference
	differences[3]++

	for _, value := range input {
		difference := value - currentJolts

		differences[difference]++

		currentJolts = value
	}

	fmt.Println(differences[1] * differences[3])
}

func partTwo(input []int) {
	// Prepend '0' to our input to count as the charger
	input = append([]int{ 0 }, input...)
	// Append our device's value (largest input value + 3)
	input = append(input, input[len(input) - 1] + 3)

	// Build a map of all possible next options for each value in our input
	options := make(map[int][]int)

	for i := 0; i < len(input); i++ {
		var next []int
		// While j exists within the input, and it isn't more than 3 values away from i
		// Then it can be considered a potential next value
		for j := i + 1; j < len(input) && input[j] - input[i] <= 3; j++ {
			next = append(next, input[j])
		}
		options[input[i]] = next
	}

	last := input[len(input) - 1]

	// Create our cache to speed up the process
	cache := make(map[int]int64)

	permutationCount := permutations(0, options, last, cache)

	fmt.Println(permutationCount)
}

func permutations(currentValue int, options map[int][]int, last int, cache map[int]int64) int64 {
	var total int64

	// Break condition for our recursion
	if currentValue == last {
		return 1
	}

	// Iterate over all options for the current value
	for _, option := range options[currentValue] {
		if cache[option] == 0 {
			total += permutations(option, options, last, cache)
		} else {
			total += cache[option]
		}
	}

	// Cache the count of permutations created for the current value
	cache[currentValue] = total

	return total
}
