package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)
	totalJoltage := 0

	for reader.Scan() {
		line := reader.Text()
		currJoltage := MaxJoltage(line, 12)
		totalJoltage += currJoltage
		fmt.Printf("Joltage for bank %s: %d\n", line, currJoltage)
	}

	fmt.Printf("Max combined joltage: %d\n", totalJoltage)
}

// Given a string, find non-contigious substring that is of length targetLength
// and contains the maximum integer value possible
//
// Ex: "818181911112111" -> 888911112111
func MaxJoltage(bank string, targetLength int) int {
	currMaxString := string(bank[0])

	for i, char := range bank[1:] {
		i++
		curr := charToInt(byte(char))

		// Given how many characters we have in our max & how many characters are left in the bank,
		// Check how many digits we can replace within our max so we can still make a total of 12 digits
		charsLeftInBank := len(bank) - i
		charsLeftToFillIn := targetLength - len(currMaxString)
		digitsWeCanReplace := min(charsLeftInBank, targetLength) - charsLeftToFillIn

		// Get the most significant digit we can replace
		startIdx := len(currMaxString) - digitsWeCanReplace
		replaceIdx := FindIndexSmallerThanDigit(curr, currMaxString, startIdx)

		// If we can replace a digit, replace and remove the preceding digits
		if replaceIdx > -1 {
			newMax := currMaxString[:replaceIdx]
			newMax += string(char)
			currMaxString = newMax
		} else if len(currMaxString) < targetLength {
			// Add string if we didn't replace it and we have more space in our number
			currMaxString += string(char)
		}
	}

	res, err := strconv.Atoi(currMaxString)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func charToInt(char byte) int {
	return int(char % 48)
}

// Given a digit, get the first index in the given string where the digit is < the input
// Starts searching at and including start. Returns -1 if all digits are >= the input
func FindIndexSmallerThanDigit(digit int, str string, start int) int {
	for i := start; i < len(str); i++ {
		if charToInt(str[i]) < digit {
			return i
		}
	}
	return -1
}
