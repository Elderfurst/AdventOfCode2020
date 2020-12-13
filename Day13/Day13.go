package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var earliestTime int

func main() {
	input := readInput()

	partOne(input)
	partTwo(input)
}

func readInput() (buses []string) {
	// Open our input file
	file, err := os.Open("Day13/Day13.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Get the first line, which is the earliest time we can depart
	scanner.Scan()

	earliestTime, _ = strconv.Atoi(scanner.Text())

	// Scan again to get the list of buses
	scanner.Scan()

	// Get the list of buses
	text := scanner.Text()

	// Split on comma to get individual buses
	buses = strings.Split(text, ",")

	return buses
}

func partOne(input []string) {
	earliestBusNumber := math.MaxInt32
	leastWaitMinutes := math.MaxInt32

	for _, bus := range input {
		// Convert our bus number to an int
		busNumber, err := strconv.Atoi(bus)

		// Skip any out of order ('x') buses
		if err != nil {
			continue
		}

		// Determine how long you'll be waiting for this bus
		remainder := earliestTime % busNumber

		waitMinutes := busNumber - remainder

		// Check to see if this is the current least amount of time you would need to wait
		if waitMinutes < leastWaitMinutes {
			leastWaitMinutes = waitMinutes
			earliestBusNumber = busNumber
		}
	}

	fmt.Println(leastWaitMinutes * earliestBusNumber)
}

func partTwo(input []string) {
	// Establish our result variable, and the amount we want to step each iteration
	currentTime, step := 0, 1

	for i, bus := range input {
		busNumber, err := strconv.Atoi(bus)
		// Skip any of the records where bus timing doesn't matter ('x')
		if err != nil {
			continue
		}

		// Loop through increments of our step value until we find one that works for the current bus
		for (currentTime + i) % busNumber != 0 {
			currentTime += step
		}

		// Increment our step
		// Multiplying because once we find a value that satisfies each bus
		// Stepping by that value every time will always be a valid time for that bus
		step *= busNumber
	}

	fmt.Println(currentTime)
}
