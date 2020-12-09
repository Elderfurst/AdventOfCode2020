package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	partOne(input)
	partTwo(input)
}

func readInput() (records []*Instruction) {
	// Open our input file
	file, err := os.Open("Day8/Day8.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records = make([]*Instruction, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Split on space, instruction is to the left value is to the right
		splitText := strings.Split(text, " ")

		convertedValue, _ := strconv.Atoi(splitText[1])

		instruction := Instruction {
			Name: splitText[0],
			Value: convertedValue,
		}

		records = append(records, &instruction)
	}

	return records
}

func partOne(instructions []*Instruction) {
	_, value := runProgram(instructions)

	fmt.Println(value)
}

func partTwo(instructions []*Instruction) {
	for _, instruction := range instructions {
		// Swap 'nop' and 'jmp' if applicable
		if instruction.Name == "nop" {
			instruction.Name = "jmp"
		} else if instruction.Name == "jmp" {
			instruction.Name = "nop"
		} else {
			continue
		}

		// Run the program with the swapped instruction
		success, value := runProgram(instructions)

		if success {
			// If we succeed then just print the value and stop executing
			fmt.Println(value)
			break
		} else {
			// Otherwise swap the instruction back
			if instruction.Name == "nop" {
				instruction.Name = "jmp"
			} else if instruction.Name == "jmp" {
				instruction.Name = "nop"
			}
		}

		// Reset 'HasExecuted' for all of our instructions after the run so it doesn't impact future runs
		for _, reset := range instructions {
			reset.HasExecuted = false
		}
	}
}

func runProgram(instructions []*Instruction) (bool, int) {
	// Instantiate our accumulator
	accumulator := 0

	for i := 0; i < len(instructions); {
		instruction := instructions[i]

		// If we've seen this instruction before, return false and the accumulator value
		if instruction.HasExecuted == true {
			return false, accumulator
		}

		// Mark this instruction as executed
		instruction.HasExecuted = true

		// Perform the expected action
		switch instruction.Name {
		case "nop":
			// Just increment our index for the next action and do nothing else
			i++
			continue
		case "acc":
			// Increment the index for the next action, and adjust the accumulator value
			i++
			accumulator += instruction.Value
		case "jmp":
			// Set the index for the desired action
			i += instruction.Value
		}
	}

	// If we make it to the end of the instructions then we've succeeded
	// return true and the accumulator value
	return true, accumulator
}

type Instruction struct {
	Name        string
	Value       int
	HasExecuted bool
}
