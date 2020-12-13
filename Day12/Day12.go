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

func readInput() (records []Instruction) {
	// Open our input file
	file, err := os.Open("Day12/Day12.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records = make([]Instruction, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Parse out the command and the value
		// The command will always be the first character
		// The value will be the rest
		command := text[:1]

		value := text[1:]

		// Convert the value to an int
		convertedValue, _ := strconv.Atoi(value)

		instruction := Instruction {
			Command: command,
			Value: convertedValue,
		}

		records = append(records, instruction)
	}

	return records
}

func partOne(input []Instruction) {
	direction := "E"

	horizontal := 0
	vertical := 0

	for _, instruction := range input {
		switch instruction.Command {
		case "F":
			switch direction {
			case "N":
				vertical += instruction.Value
			case "S":
				vertical -= instruction.Value
			case "E":
				horizontal += instruction.Value
			case "W":
				horizontal -= instruction.Value
			}
		case "N":
			vertical += instruction.Value
		case "S":
			vertical -= instruction.Value
		case "E":
			horizontal += instruction.Value
		case "W":
			horizontal -= instruction.Value
		case "L":
			switch instruction.Value {
			case 90:
				switch direction {
				case "N":
					direction = "W"
				case "S":
					direction = "E"
				case "E":
					direction = "N"
				case "W":
					direction = "S"
				}
			case 180:
				switch direction {
				case "N":
					direction = "S"
				case "S":
					direction = "N"
				case "E":
					direction = "W"
				case "W":
					direction = "E"
				}
			case 270:
				switch direction {
				case "N":
					direction = "E"
				case "S":
					direction = "W"
				case "E":
					direction = "S"
				case "W":
					direction = "N"
				}
			}
		case "R":
			switch instruction.Value {
			case 90:
				switch direction {
				case "N":
					direction = "E"
				case "S":
					direction = "W"
				case "E":
					direction = "S"
				case "W":
					direction = "N"
				}
			case 180:
				switch direction {
				case "N":
					direction = "S"
				case "S":
					direction = "N"
				case "E":
					direction = "W"
				case "W":
					direction = "E"
				}
			case 270:
				switch direction {
				case "N":
					direction = "W"
				case "S":
					direction = "E"
				case "E":
					direction = "N"
				case "W":
					direction = "S"
				}
			}
		}
	}

	manhattanDistance := absoluteValue(vertical) + absoluteValue(horizontal)

	fmt.Println(manhattanDistance)
}

func partTwo(input []Instruction) {
	shipHorizontal := 0
	shipVertical := 0

	waypointHorizontal := 10
	waypointVertical := 1

	for _, instruction := range input {
		switch instruction.Command {
		case "F":
			for i := 0; i < instruction.Value; i++ {
				shipHorizontal += waypointHorizontal
				shipVertical += waypointVertical
			}
		case "N":
			waypointVertical += instruction.Value
		case "S":
			waypointVertical -= instruction.Value
		case "E":
			waypointHorizontal += instruction.Value
		case "W":
			waypointHorizontal -= instruction.Value
		case "L":
			switch instruction.Value {
			case 90:
				waypointHorizontal, waypointVertical = rotateLeft(waypointHorizontal, waypointVertical)
			case 180:
				waypointHorizontal, waypointVertical = rotateLeft(waypointHorizontal, waypointVertical)
				waypointHorizontal, waypointVertical = rotateLeft(waypointHorizontal, waypointVertical)
			case 270:
				waypointHorizontal, waypointVertical = rotateRight(waypointHorizontal, waypointVertical)
			}
		case "R":
			switch instruction.Value {
			case 90:
				waypointHorizontal, waypointVertical = rotateRight(waypointHorizontal, waypointVertical)
			case 180:
				waypointHorizontal, waypointVertical = rotateRight(waypointHorizontal, waypointVertical)
				waypointHorizontal, waypointVertical = rotateRight(waypointHorizontal, waypointVertical)
			case 270:
				waypointHorizontal, waypointVertical = rotateLeft(waypointHorizontal, waypointVertical)
			}
		}
	}

	manhattanDistance := absoluteValue(shipVertical) + absoluteValue(shipHorizontal)

	fmt.Println(manhattanDistance)
}

func absoluteValue(value int) int {
	if value < 0 {
		value = -value
	}

	return value
}

func rotateRight(x, y int) (newX, newY int) {
	newX = y
	newY = -x

	return newX, newY
}

func rotateLeft(x, y int) (newX, newY int) {
	newX = -y
	newY = x

	return newX, newY
}

type Instruction struct {
	Command string
	Value   int
}
