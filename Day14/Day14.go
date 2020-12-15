package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	partOne(input)
	partTwo(input)
}

func readInput() (lines []string) {
	// Open our input file
	file, err := os.Open("Day14/Day14.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines = make([]string, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		lines = append(lines, text)
	}

	return lines
}

func partOne(input []string) {
	currentMask := ""

	memory := make(map[int]string, 0)

	for _, line := range input {
		// Split on ' = ' to include the spaces surrounding the equal sign, so we don't have to trim after
		splitLine := strings.Split(line, " = ")

		command := splitLine[0]
		value := splitLine[1]

		// If we get a mask line, then just set the mask and move on
		if command == "mask" {
			currentMask = value
			continue
		}

		// Pull the memory index out of the first part of the line
		memoryIndex := getMemoryIndex(command)

		// Convert the second part of the line to an int
		parsedValue, _ := strconv.Atoi(value)

		// Convert the above parsed value to a string containing its binary representation
		binaryValue := expandNumber(parsedValue)

		// Split the binary number into an array for easy adjusting
		splitBinary := strings.Split(binaryValue, "")

		// Split our current mask into an array for easy iterating
		splitMask := strings.Split(currentMask, "")

		// Iterate over the current mask and adjust any values necessary
		for i, operator := range splitMask {
			if operator == "X" {
				continue
			}

			splitBinary[i] = operator
		}

		// Rejoin our binary string after adjustments for storage
		memory[memoryIndex] = strings.Join(splitBinary, "")
	}

	var sum int

	// Convert all of our stored binary numbers back to integers and sum them
	for _, memoryValue := range memory {
		contractedValue := contractNumber(memoryValue)

		sum += contractedValue
	}

	fmt.Println(sum)
}

func partTwo(input []string) {
	currentMask := ""

	memory := make(map[string]int64, 0)

	for _, line := range input {
		// Split on ' = ' to include the spaces surrounding the equal sign, so we don't have to trim after
		splitLine := strings.Split(line, " = ")

		command := splitLine[0]
		value := splitLine[1]

		// If we get a mask line, then just set the mask and move on
		if command == "mask" {
			currentMask = value
			continue
		}

		// Pull the memory index out of the first part of the line
		memoryIndex := getMemoryIndex(command)

		// Split our current mask into an array for easy iterating
		splitMask := strings.Split(currentMask, "")

		binaryMemoryIndex := expandNumber(memoryIndex)

		splitBinaryMemoryIndex := strings.Split(binaryMemoryIndex, "")

		// Iterate over the current mask and adjust any values necessary
		for i, operator := range splitMask {
			switch operator {
			case "0":
				continue
			case "1":
				splitBinaryMemoryIndex[i] = "1"
			case "X":
				splitBinaryMemoryIndex[i] = "X"
			}
		}

		duplicateAddresses := getPossibleMemoryAddresses(0, strings.Join(splitBinaryMemoryIndex, ""))

		memoryAddresses := removeDuplicateAddresses(duplicateAddresses)

		// Convert the second part of the line to an int
		parsedValue, _ := strconv.ParseInt(value, 10, 64)

		for _, memoryAddress := range memoryAddresses {
			// Rejoin our binary string after adjustments for storage
			memory[memoryAddress] = parsedValue
		}
	}

	var sum int64

	// Sum everything left in memory
	for _, memoryValue := range memory {
		sum += memoryValue
	}

	fmt.Println(sum)
}

func getMemoryIndex(line string) (index int) {
	memoryIndex := regexp.MustCompile("^mem\\[([0-9]+)]$")

	regexValues := memoryIndex.FindStringSubmatch(line)

	// The first (0 index) result here is the complete match
	// The second (1 index) result is the capture group with just the index
	parsedValue, _ := strconv.Atoi(regexValues[1])

	return parsedValue
}

func expandNumber(number int) string {
	// Format our number into its binary representation and pad it out to 36 digits with 0s
	formatted := strconv.FormatInt(int64(number), 2)

	return fmt.Sprintf("%036s", formatted)
}

func contractNumber(value string) int {
	// Convert back from a binary number to a base-10 number
	convertedValue, _ := strconv.ParseInt(value, 2, 0)

	return int(convertedValue)
}

func getPossibleMemoryAddresses(index int, memoryAddress string) (memoryAddresses []string) {
	splitAddress := strings.Split(memoryAddress, "")

	tempAddress := memoryAddress

	xIndex := strings.Index(tempAddress, "X")

	// If there are no more 'X's in the address, then append the current address and return
	if xIndex < 0 {
		memoryAddresses = append(memoryAddresses, tempAddress)

		return memoryAddresses
	}

	splitTemp := strings.Split(tempAddress, "")

	for i, value := range splitAddress {
		if value == "X" {
			splitTemp[i] = "0"

			memoryAddresses = append(memoryAddresses, getPossibleMemoryAddresses(i, strings.Join(splitTemp, ""))...)

			splitTemp[i] = "1"

			memoryAddresses = append(memoryAddresses, getPossibleMemoryAddresses(i, strings.Join(splitTemp, ""))...)
		}
	}

	return memoryAddresses
}

func removeDuplicateAddresses(duplicates []string) (memoryAddresses []string) {
	keys := make(map[string]bool)

	for _, address := range duplicates {
		_, exists := keys[address]

		if !exists {
			keys[address] = true

			memoryAddresses = append(memoryAddresses, address)
		}
	}

	return memoryAddresses
}
