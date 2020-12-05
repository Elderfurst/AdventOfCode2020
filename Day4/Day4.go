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

func readInput() []Passport {
	// Open our input file
	file, err := os.Open("Day4/Day4.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var records []Passport

	scanner := bufio.NewScanner(file)

	currentPassport := new(Passport)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		// If we get an empty line then we can consider the latest passport complete and create a new one
		if text == "" {
			records = append(records, *currentPassport)
			currentPassport = new(Passport)
			continue
		}

		// Parse whatever fields are present into the current passport
		line := strings.Split(text, " ")

		for _, field := range line {
			splitField := strings.Split(field, ":")

			key := splitField[0]
			value := splitField[1]

			switch key {
			case "byr":
				currentPassport.birthYear, _ = strconv.Atoi(value)
			case "iyr":
				currentPassport.issueYear, _ = strconv.Atoi(value)
			case "eyr":
				currentPassport.expirationYear, _ = strconv.Atoi(value)
			case "hgt":
				currentPassport.height = value
			case "hcl":
				currentPassport.hairColor = value
			case "ecl":
				currentPassport.eyeColor = value
			case "pid":
				currentPassport.passportId = value
			case "cid":
				currentPassport.countryId, _ = strconv.Atoi(value)
			}
		}
	}

	records = append(records, *currentPassport)

	return records
}

func partOne(input []Passport) {
	validPassports := 0

	// Check to make sure none of the required fields are left blank
	// The only field not required is country id (cid)
	for _, passport := range input {
		if passport.birthYear != 0 &&
			passport.issueYear != 0 &&
			passport.expirationYear != 0 &&
			passport.height != "" &&
			passport.hairColor != "" &&
			passport.eyeColor != "" &&
			passport.passportId != "" {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func partTwo(input []Passport) {
	validPassports := 0

	// Check to make sure all fields are present and also valid
	for _, passport := range input {
		if validatePassportId(passport.passportId) &&
			validateBirthYear(passport.birthYear) &&
			validateIssueYear(passport.issueYear) &&
			validateExpirationYear(passport.expirationYear) &&
			validateHeight(passport.height) &&
			validateHairColor(passport.hairColor) &&
			validateEyeColor(passport.eyeColor) &&
			validatePassportId(passport.passportId) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func validateBirthYear (birthYear int) bool {
	return birthYear != 0 && birthYear >= 1920 && birthYear <= 2002
}

func validateIssueYear (issueYear int) bool {
	return issueYear != 0 && issueYear >= 2010 && issueYear <= 2020
}

func validateExpirationYear (expirationYear int) bool {
	return expirationYear != 0 && expirationYear >= 2020 && expirationYear <= 2030
}

func validateHeight (height string) bool {
	if strings.Contains(height, "cm") {
		centimeters := strings.Replace(height, "cm", "", -1)

		convertedCentimeters, _ := strconv.Atoi(centimeters)

		return convertedCentimeters >= 150 && convertedCentimeters <= 193
	} else if strings.Contains(height, "in") {
		inches := strings.Replace(height, "in", "", -1)

		convertedInches, _ := strconv.Atoi(inches)

		return convertedInches >= 59 && convertedInches <= 76
	} else {
		return false
	}
}

func validateHairColor (hairColor string) bool {
	validHairColor := regexp.MustCompile("^#[a-f0-9]{6}$")

	return validHairColor.MatchString(hairColor)
}

func validateEyeColor (eyeColor string) bool {
	return eyeColor == "amb" ||
		eyeColor == "blu" ||
		eyeColor == "brn" ||
		eyeColor == "gry" ||
		eyeColor == "grn" ||
		eyeColor == "hzl" ||
		eyeColor == "oth"
}

func validatePassportId (passportId string) bool {
	validPassportId := regexp.MustCompile("^[0-9]{9}$")

	return validPassportId.MatchString(passportId)
}

type Passport struct {
	birthYear      int
	issueYear 	   int
	expirationYear int
	height         string
	hairColor      string
	eyeColor       string
	passportId     string
	countryId      int
}
