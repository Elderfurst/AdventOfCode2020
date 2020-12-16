package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, tickets := readInput()

	validTickets := partOne(rules, tickets)
	partTwo(rules, validTickets, tickets[0])
}

func readInput() (rules []Rule, tickets [][]int) {
	// Open our input file
	file, err := os.Open("Day16/Day16.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules = make([]Rule, 0)
	tickets = make([][]int, 0)

	mode := "rules"

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Skip empty lines
		if text == "" {
			continue
		}

		// Change the mode to tickets when necessary
		if text == "your ticket:" || text == "nearby tickets:" {
			mode = "tickets"
			continue
		}

		switch mode {
		case "rules":
			// Split on colon to get the name
			splitText := strings.Split(text, ":")

			secondSection := strings.TrimSpace(splitText[1])

			// Split the second section on space to get the various min and max values
			splitValues := strings.Split(secondSection, " ")

			firstSet := splitValues[0]

			// split our values on dash to get the mins and maxes
			firstValues := strings.Split(firstSet, "-")

			firstMin, _ := strconv.Atoi(firstValues[0])
			firstMax, _ := strconv.Atoi(firstValues[1])

			secondSet := splitValues[2]

			secondValues := strings.Split(secondSet, "-")

			secondMin, _ := strconv.Atoi(secondValues[0])
			secondMax, _ := strconv.Atoi(secondValues[1])


			rule := Rule {
				Name: splitText[0],
				FirstMinValue: firstMin,
				FirstMaxValue: firstMax,
				SecondMinValue: secondMin,
				SecondMaxValue: secondMax,
			}

			rules = append(rules, rule)

		case "tickets":
			ticket := make([]int, 0)

			ticketValues := strings.Split(text, ",")

			for _, ticketValue := range ticketValues {
				parsedValue, _ := strconv.Atoi(ticketValue)

				ticket = append(ticket, parsedValue)
			}

			tickets = append(tickets, ticket)
		}
	}

	return rules, tickets
}

func partOne(rules []Rule, tickets [][]int) (validTickets [][]int) {
	validTickets = make([][]int, 0)

	scanningErrorRate := 0

	// Skip our first ticket because that's ours
	for i := 1; i < len(tickets); i++ {
		pendingTicket := tickets[i]

		validTicket := true

		// Check each number on a ticket to make sure they're all valid
		for _, value := range pendingTicket {
			ruleBreaker := true

			for _, rule := range rules {
				// Check to see if this value satisfies at least one rule
				if (value >= rule.FirstMinValue && value <= rule.FirstMaxValue) ||
					(value >= rule.SecondMinValue && value <= rule.SecondMaxValue) {
					ruleBreaker = false
				}
			}

			if ruleBreaker {
				scanningErrorRate += value

				validTicket = false
			}
		}

		if validTicket {
			validTickets = append(validTickets, pendingTicket)
		}
	}

	fmt.Println(scanningErrorRate)

	return validTickets
}

func partTwo(rules []Rule, tickets [][]int, yourTicket []int) {
	possiblePositionsMap := make(map[string][]int, 0)

	// Determine the position of the field for each rule
	for _, rule := range rules {
		possiblePositions := make([]int, 0)

		// Check each ticket to see
		for i, ticket := range tickets {
			// Build our list of possible positions using the first ticket
			// Then whittle it down using the rest of the tickets
			if i == 0 {
				for position, value := range ticket {
					// Check to see if this value satisfies the current rule
					if (value >= rule.FirstMinValue && value <= rule.FirstMaxValue) ||
						(value >= rule.SecondMinValue && value <= rule.SecondMaxValue) {
						possiblePositions = append(possiblePositions, position)
					}
				}
			} else {
				newPossiblePositions := make([]int, 0)

				// Iterate over the possible positions and remove any if they don't make sense
				for _, position := range possiblePositions {
					value := ticket[position]
					// Check to see if this value satisfies the current rule
					if (value >= rule.FirstMinValue && value <= rule.FirstMaxValue) ||
						(value >= rule.SecondMinValue && value <= rule.SecondMaxValue) {
						newPossiblePositions = append(newPossiblePositions, position)
					}
				}

				possiblePositions = newPossiblePositions
			}
		}

		possiblePositionsMap[rule.Name] = possiblePositions
	}

	hardPositions := make(map[string]int, 0)

	// We need to loop until every field has only a single possible value
	allSingleValues := false

	for allSingleValues == false {
		allSingleValues = true

		for rule, rulePositions := range possiblePositionsMap {
			// If this rule is left with a single position, add it to our map and move on
			if len(rulePositions) == 1 {
				hardPositions[rule] = rulePositions[0]
			} else {
				// Only mess around with the possible positions if there are hard positions that could be removed
				if len(hardPositions) > 0 {
					newRulePositions := make([]int, 0)

					// Loop through all current possible positions and remove any records
					// That match an existing 'hard position'
					// A position can only be occupied by at most one field, so it is no longer possible for this field
					for _, possiblePosition := range rulePositions {
						shouldKeep := true

						for _, hardPosition := range hardPositions {
							if hardPosition == possiblePosition {
								shouldKeep = false
							}
						}

						if shouldKeep {
							newRulePositions = append(newRulePositions, possiblePosition)
						}
					}

					// Reset the possible positions to our newly trimmed list
					possiblePositionsMap[rule] = newRulePositions

					// If the new list has a single record, then just go ahead and add to the hard position map
					if len(newRulePositions) == 1 {
						hardPositions[rule] = newRulePositions[0]
					} else {
						// If there are still more than one possible position then we need to keep looping
						allSingleValues = false
					}
				} else {
					allSingleValues = false
				}
			}
		}
	}

	// Multiply all fields that start with 'departure' together
	answer := 1

	for rule, position := range hardPositions {
		if strings.HasPrefix(rule, "departure") {
			answer *= yourTicket[position]
		}
	}

	fmt.Println(answer)
}

type Rule struct {
	Name           string
	FirstMinValue  int
	FirstMaxValue  int
	SecondMinValue int
	SecondMaxValue int
}
