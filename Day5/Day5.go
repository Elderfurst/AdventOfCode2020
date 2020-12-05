package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := readInput()

	seats := getSeatIds(input)

	partOne(seats)
	partTwo(seats)
}

func readInput() []string {
	// Open our input file
	file, err := os.Open("Day5/Day5.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var records []string

	scanner := bufio.NewScanner(file)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		records = append(records, text)
	}

	return records
}

func getSeatIds(input []string) (seats []int) {
	for _, record := range input {
		rowDirections := strings.Split(record[0:7], "")
		columnDirections := strings.Split(record[7:10], "")

		minRow := 0
		maxRow := 127

		for _, rowDirection := range rowDirections {
			switch rowDirection {
			case "F":
				maxRow = minRow + (((maxRow - 1) - minRow) / 2)
			case "B":
				minRow = maxRow - (((maxRow - 1) - minRow) / 2)
			}
		}

		minColumn := 0
		maxColumn := 7

		for _, columnDirection := range columnDirections {
			switch columnDirection {
			case "L":
				maxColumn = minColumn + (((maxColumn - 1) - minColumn) / 2)
			case "R":
				minColumn = maxColumn - (((maxColumn - 1) - minColumn) / 2)
			}
		}

		seatId := minRow * 8 + minColumn

		seats = append(seats, seatId)
	}

	return seats
}

func partOne(seats []int) {
	maxSeatId := 0

	// find the highest seat id in our set
	for _, seat := range seats {
		if seat > maxSeatId {
			maxSeatId = seat
		}
	}

	fmt.Println(maxSeatId)
}

func partTwo(seats []int) {
	// sort the seats so they are in ascending order
	sort.Ints(seats)

	// The first seat is the minimum value, so all other seats should be equal to their position plus this value
	seatOffset := seats[0]

	for i, seat := range seats {
		// If we find a seat that doesn't match the index plus the offset
		// then there is a missing value, which is our seat
		if seat != i + seatOffset {
			fmt.Println(i + seatOffset)
			break
		}
	}
}
