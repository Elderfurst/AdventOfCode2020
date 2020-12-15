package main

import "fmt"

func main() {
	input := []int { 0, 3, 1, 6, 7, 5 }

	speakNumbers(input, 2020)
	speakNumbers(input, 30000000)
}

func speakNumbers(input []int, turnCount int) {
	pastTurns := make (map[int][]int, 0)

	turn := 0

	currentNumber := 0

	// Initialize our past turns with our input values
	for i, value := range input {
		if i > 0 {
			previousNumber := input[i - 1]

			pastTurns[previousNumber] = append(pastTurns[previousNumber], turn)
		}

		currentNumber = value

		turn++
	}

	for ; turn < turnCount; turn++ {
		// Append our current number at the current turn
		pastTurns[currentNumber] = append(pastTurns[currentNumber], turn)

		// If this number hasn't been spoken before this turn, then our next number is 0 no matter what
		if len(pastTurns[currentNumber]) == 1 {
			currentNumber = 0
		} else {
			// If the number has been spoken before, then we need to find the age difference between the last two times
			previousTurns := pastTurns[currentNumber]

			previousTurn := previousTurns[len(previousTurns) - 1]

			turnBeforeThat := previousTurns[len(previousTurns) - 2]

			currentNumber = previousTurn - turnBeforeThat
		}
	}

	fmt.Println(currentNumber)
}
