package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	groups := buildGroups(input)

	partOne(groups)
	partTwo(groups)
}

func readInput() (records [][]string) {
	// Open our input file
	file, err := os.Open("Day6/Day6.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Split on nothing so we get each character as a single entry in the array
		splitText := strings.Split(text, "")

		records = append(records, splitText)
	}

	return records
}

func buildGroups(input [][]string) (groups []Group) {
	currentGroup := new(Group)
	currentGroup.Answers = make(map[string]int, 0)

	// Loop over each line of our input
	// Empty lines denote a change in group
	for _, answer := range input {
		if len(answer) == 0 {
			groups = append(groups, *currentGroup)

			currentGroup = new(Group)
			currentGroup.Answers = make(map[string]int, 0)
			continue
		}
		// Increment the count, which represents the number of people in the group
		currentGroup.Count++

		// Loop over the values in each line, these represent individual 'questions'
		for _, question := range answer {
			_, exists := currentGroup.Answers[question]

			if !exists {
				currentGroup.Answers[question] = 1
			} else {
				currentGroup.Answers[question]++
			}
		}
	}

	// Append the last group since it doesn't happen automatically due to our input
	groups = append(groups, *currentGroup)

	return groups
}

func partOne(groups []Group) {
	sumOfAnswers := 0

	// Next we just need to count the number of unique questions each group answered and add them all up
	for _, group := range groups {
		for _, _ = range group.Answers {
			sumOfAnswers++
		}
	}

	fmt.Println(sumOfAnswers)
}

func partTwo(groups []Group) {
	sumOfAnswers := 0

	// Next we just need to count the number of unique questions each group answered and add them all up
	for _, group := range groups {
		for _, value := range group.Answers {
			if value == group.Count {
				sumOfAnswers++
			}
		}
	}

	fmt.Println(sumOfAnswers)
}

type Group struct {
	Count   int
	Answers map[string]int
}
