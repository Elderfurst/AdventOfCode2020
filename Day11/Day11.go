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

func readInput() (records [][]string) {
	// Open our input file
	file, err := os.Open("Day11/Day11.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records = make([][]string, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Split on nothing to get each character in an array
		splitRow := strings.Split(text, "")

		// Append our split row to our records array
		records = append(records, splitRow)
	}

	return records
}

func partOne(input [][]string) {
	for seatsChanged := true; seatsChanged == true; {
		seatsChanged = false

		newValues := make([][]string, len(input))

		for i, row := range input {
			newValues[i] = make([]string, len(row))
			for j, _ := range row {
				switch input[i][j] {
				case ".":
					newValues[i][j] = "."
					continue
				case "L":
					newValues[i][j] = "L"
					// Check all adjacent seats clockwise from up and left
					upLeftX := j - 1
					upLeftY := i - 1

					if upLeftX >= 0 && upLeftY >= 0 {
						upLeftValue := input[upLeftY][upLeftX]

						// If any seat is occupied around an empty seat then it won't become occupied
						if upLeftValue == "#" {
							continue
						}
					}

					upX := j
					upY := i - 1

					if upY >= 0 {
						aboveValue := input[upY][upX]

						if aboveValue == "#" {
							continue
						}
					}

					upRightX := j + 1
					upRightY := i - 1

					if upRightX < len(row) && upRightY >= 0 {
						upRightValue := input[upRightY][upRightX]

						if upRightValue == "#" {
							continue
						}
					}

					rightX := j + 1
					rightY := i

					if rightX < len(row) {
						rightValue := input[rightY][rightX]

						if rightValue == "#" {
							continue
						}
					}

					downRightX := j + 1
					downRightY := i + 1

					if downRightX < len(row) && downRightY < len(input) {
						downRightValue := input[downRightY][downRightX]

						if downRightValue == "#" {
							continue
						}
					}

					downX := j
					downY := i + 1

					if downY < len(input) {
						downValue := input[downY][downX]

						if downValue == "#" {
							continue
						}
					}

					downLeftX := j - 1
					downLeftY := i + 1

					if downLeftX >= 0 && downLeftY < len(input) {
						downLeftValue := input[downLeftY][downLeftX]

						if downLeftValue == "#" {
							continue
						}
					}

					leftX := j - 1
					leftY := i

					if leftX >= 0 {
						leftValue := input[leftY][leftX]

						if leftValue == "#" {
							continue
						}
					}

					newValues[i][j] = "#"
					seatsChanged = true
				case "#":
					newValues[i][j] = "#"
					adjacentOccupiedSeats := 0
					// Check all adjacent seats clockwise from up and left
					upLeftX := j - 1
					upLeftY := i - 1

					if upLeftX >= 0 && upLeftY >= 0 {
						upLeftValue := input[upLeftY][upLeftX]

						if upLeftValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					upX := j
					upY := i - 1

					if upY >= 0 {
						aboveValue := input[upY][upX]

						if aboveValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					upRightX := j + 1
					upRightY := i - 1

					if upRightX < len(row) && upRightY >= 0 {
						upRightValue := input[upRightY][upRightX]

						if upRightValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					rightX := j + 1
					rightY := i

					if rightX < len(row) {
						rightValue := input[rightY][rightX]

						if rightValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					downRightX := j + 1
					downRightY := i + 1

					if downRightX < len(row) && downRightY < len(input) {
						downRightValue := input[downRightY][downRightX]

						if downRightValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					downX := j
					downY := i + 1

					if downY < len(input) {
						downValue := input[downY][downX]

						if downValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					downLeftX := j - 1
					downLeftY := i + 1

					if downLeftX >= 0 && downLeftY < len(input) {
						downLeftValue := input[downLeftY][downLeftX]

						if downLeftValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					leftX := j - 1
					leftY := i

					if leftX >= 0 {
						leftValue := input[leftY][leftX]

						if leftValue == "#" {
							adjacentOccupiedSeats++
						}
					}

					if adjacentOccupiedSeats >= 4 {
						newValues[i][j] = "L"
						seatsChanged = true
					}
				}
			}
		}

		input = newValues
	}

	occupiedSeats := 0

	for i, row := range input {
		for j, _ := range row {
			if input[i][j] == "#" {
				occupiedSeats++
			}
		}
	}

	fmt.Println(occupiedSeats)
}

func partTwo(input [][]string) {
	for seatsChanged := true; seatsChanged == true; {
		seatsChanged = false

		newValues := make([][]string, len(input))

		for i, row := range input {
			newValues[i] = make([]string, len(row))
			for j, _ := range row {
				switch input[i][j] {
				case ".":
					newValues[i][j] = "."
					continue
				case "L":
					newValues[i][j] = "L"
					// Check all adjacent seats clockwise from up and left
					upLeftX := j
					upLeftY := i
					upLeftValue := "."

					for ; upLeftX > 0 && upLeftY > 0; {
						upLeftX -= 1
						upLeftY -= 1

						upLeftValue = input[upLeftY][upLeftX]

						if upLeftValue == "L" || upLeftValue == "#" {
							break
						}
					}

					// If any seat is occupied around an empty seat then it won't become occupied
					if upLeftValue == "#" {
						continue
					}

					upX := j
					upY := i

					upValue := "."

					for ; upY > 0; {
						upY -= 1

						upValue = input[upY][upX]

						if upValue == "L" || upValue == "#" {
							break
						}
					}

					if upValue == "#" {
						continue
					}

					upRightX := j
					upRightY := i

					upRightValue := "."

					for ; upRightX < len(row) - 1 && upRightY > 0; {
						upRightX += 1
						upRightY -= 1

						upRightValue = input[upRightY][upRightX]

						if upRightValue == "L" || upRightValue == "#" {
							break
						}
					}

					if upRightValue == "#" {
						continue
					}

					rightX := j
					rightY := i

					rightValue := "."

					for ; rightX < len(row) - 1; {
						rightX += 1

						rightValue = input[rightY][rightX]

						if rightValue == "L" || rightValue == "#" {
							break
						}
					}

					if rightValue == "#" {
						continue
					}

					downRightX := j
					downRightY := i

					downRightValue := "."

					for ; downRightX < len(row) - 1 && downRightY < len(input) - 1; {
						downRightX += 1
						downRightY += 1

						downRightValue = input[downRightY][downRightX]

						if downRightValue == "L" || downRightValue == "#" {
							break
						}
					}

					if downRightValue == "#" {
						continue
					}

					downX := j
					downY := i

					downValue := "."

					for ; downY < len(input) - 1; {
						downY += 1

						downValue = input[downY][downX]

						if downValue == "L" || downValue == "#" {
							break
						}
					}

					if downValue == "#" {
						continue
					}

					downLeftX := j
					downLeftY := i

					downLeftValue := "."

					for ; downLeftX > 0 && downLeftY < len(input) - 1; {
						downLeftX -= 1
						downLeftY += 1

						downLeftValue = input[downLeftY][downLeftX]

						if downLeftValue == "L" || downLeftValue == "#" {
							break
						}
					}

					if downLeftValue == "#" {
						continue
					}

					leftX := j
					leftY := i

					leftValue := "."

					for ; leftX > 0; {
						leftX -= 1

						leftValue = input[leftY][leftX]

						if leftValue == "L" || leftValue == "#" {
							break
						}
					}

					if leftValue == "#" {
						continue
					}

					newValues[i][j] = "#"
					seatsChanged = true
				case "#":
					newValues[i][j] = "#"
					adjacentOccupiedSeats := 0

					// Check all adjacent seats clockwise from up and left
					upLeftX := j
					upLeftY := i
					upLeftValue := "."

					for ; upLeftX > 0 && upLeftY > 0; {
						upLeftX -= 1
						upLeftY -= 1

						upLeftValue = input[upLeftY][upLeftX]

						if upLeftValue == "L" || upLeftValue == "#" {
							break
						}
					}

					// If any seat is occupied around an empty seat then it won't become occupied
					if upLeftValue == "#" {
						adjacentOccupiedSeats++
					}

					upX := j
					upY := i

					upValue := "."

					for ; upY > 0; {
						upY -= 1

						upValue = input[upY][upX]

						if upValue == "L" || upValue == "#" {
							break
						}
					}

					if upValue == "#" {
						adjacentOccupiedSeats++
					}

					upRightX := j
					upRightY := i

					upRightValue := "."

					for ; upRightX < len(row) - 1 && upRightY > 0; {
						upRightX += 1
						upRightY -= 1

						upRightValue = input[upRightY][upRightX]

						if upRightValue == "L" || upRightValue == "#" {
							break
						}
					}

					if upRightValue == "#" {
						adjacentOccupiedSeats++
					}

					rightX := j
					rightY := i

					rightValue := "."

					for ; rightX < len(row) - 1; {
						rightX += 1

						rightValue = input[rightY][rightX]

						if rightValue == "L" || rightValue == "#" {
							break
						}
					}

					if rightValue == "#" {
						adjacentOccupiedSeats++
					}

					downRightX := j
					downRightY := i

					downRightValue := "."

					for ; downRightX < len(row) - 1 && downRightY < len(input) - 1; {
						downRightX += 1
						downRightY += 1

						downRightValue = input[downRightY][downRightX]

						if downRightValue == "L" || downRightValue == "#" {
							break
						}
					}

					if downRightValue == "#" {
						adjacentOccupiedSeats++
					}

					downX := j
					downY := i

					downValue := "."

					for ; downY < len(input) - 1; {
						downY += 1

						downValue = input[downY][downX]

						if downValue == "L" || downValue == "#" {
							break
						}
					}

					if downValue == "#" {
						adjacentOccupiedSeats++
					}

					downLeftX := j
					downLeftY := i

					downLeftValue := "."

					for ; downLeftX > 0 && downLeftY < len(input) - 1; {
						downLeftX -= 1
						downLeftY += 1

						downLeftValue = input[downLeftY][downLeftX]

						if downLeftValue == "L" || downLeftValue == "#" {
							break
						}
					}

					if downLeftValue == "#" {
						adjacentOccupiedSeats++
					}

					leftX := j
					leftY := i

					leftValue := "."

					for ; leftX > 0; {
						leftX -= 1

						leftValue = input[leftY][leftX]

						if leftValue == "L" || leftValue == "#" {
							break
						}
					}

					if leftValue == "#" {
						adjacentOccupiedSeats++
					}

					if adjacentOccupiedSeats >= 5 {
						newValues[i][j] = "L"
						seatsChanged = true
					}
				}
			}
		}

		input = newValues
	}

	occupiedSeats := 0

	for i, row := range input {
		for j, _ := range row {
			if input[i][j] == "#" {
				occupiedSeats++
			}
		}
	}

	fmt.Println(occupiedSeats)
}
