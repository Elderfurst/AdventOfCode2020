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

func readInput() (records map[string]*Bag) {
	// Open our input file
	file, err := os.Open("Day7/Day7.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records = make(map[string]*Bag, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// Split on contains, the left side will be the bag the rule applies to, the right side will be the rule itself
		splitText := strings.Split(text, " contain ")

		mainBag := splitText[0]

		// Remove the trailing 's' so all bags are indexed singularly
		if string(mainBag[len(mainBag) - 1]) == "s" {
			mainBag = mainBag[:len(mainBag) - 1]
		}

		// Check to see if a bag was created previously and use it if so
		bag, exists := records[mainBag]

		// If the bag doesn't exist already then create it and add it to the map
		if !exists {
			newBag := Bag {
				Name: mainBag,
			}

			bag = &newBag

			records[bag.Name] = bag
		}

		children := strings.Split(splitText[1], ",")

		for _, child := range children {
			// Remove any surrounding whitespace or the trailing period
			parsedChild := strings.TrimSpace(child)
			parsedChild = strings.Trim(parsedChild, ".")

			// Split on space to pull the quantity out cleanly
			splitChild := strings.Split(parsedChild, " ")

			// Assign the quantity and build the name back up
			childQuantity, _ := strconv.Atoi(splitChild[0])
			childName := strings.Join(splitChild[1:], " ")

			// Remove the trailing 's' so all bags are indexed singularly
			if string(childName[len(childName) - 1]) == "s" {
				childName = childName[:len(childName) - 1]
			}

			// Check to see if this child already exists as a bag in our map
			existingBag, exists := records[childName]

			if exists {
				// If it does exist then use the same object for this rule
				bagRule := BagRule {
					Bag: existingBag,
					Quantity: childQuantity,
				}

				bag.BagRules = append(bag.BagRules, &bagRule)
			} else {
				// If it doesn't exist then we create a new bag, add it to the map, and use it in the rule
				childBag := Bag {
					Name: childName,
				}

				records[childBag.Name] = &childBag

				childBagRule := BagRule {
					Bag: &childBag,
					Quantity: childQuantity,
				}

				bag.BagRules = append(bag.BagRules, &childBagRule)
			}
		}
	}

	return records
}

func partOne(bags map[string]*Bag) {
	totalBags := 0

	for _, bag := range bags {
		// Check to see if there is any path from our current bag to the bag we want
		bagPath := findBag(*bag, "shiny gold bag", 1)

		// If such a path exists, then this counts
		if len(bagPath) > 0 {
			totalBags++
		}
	}

	fmt.Println(totalBags)
}

func partTwo(bags map[string]*Bag) {
	shinyGoldBag := bags["shiny gold bag"]

	bagCount := countBags(*shinyGoldBag)

	fmt.Println(bagCount)
}

func countBags(bag Bag) int {
	bagCount := 0

	for _, bagRule := range bag.BagRules {
		quantity := bagRule.Quantity
		// Since there are 'x' of each of these bags we need to multiple the bag count by that number
		bagCount += quantity + (quantity * countBags(*bagRule.Bag))
	}

	return bagCount
}

func findBag(bag Bag, bagToFind string, quantity int) (path []string) {
	for _, bagRule := range bag.BagRules {
		if bagRule.Bag.Name == bagToFind &&
			bagRule.Quantity >= quantity {
			// If the current bag has a rule that it can carry our desired bag, build the path
			path = append(path, bag.Name)
			break
		} else {
			// If not, try to build the path using this rule's bag
			path = append(path, findBag(*bagRule.Bag, bagToFind, quantity)...)
		}
	}

	return path
}

type Bag struct {
	Name     string
	BagRules []*BagRule
}

type BagRule struct {
	Bag      *Bag
	Quantity int
}
