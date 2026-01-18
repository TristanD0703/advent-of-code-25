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

func MaxJoltage(bank string, targetLength int) int {
	currMaxString := string(bank[0])
	// currMaxLeft, currMaxRight, leftIndex := charToInt(bank[0]), charToInt(bank[1]), 0

	for i, char := range bank[1:] {
		i++
		curr := charToInt(byte(char))

		charsLeftInBank := len(bank) - i
		charsLeftToFillIn := targetLength - len(currMaxString)
		digitsWeCanReplace := min(charsLeftInBank, targetLength) - charsLeftToFillIn
		replaced := false

		if digitsWeCanReplace > 0 {
			startIdx := len(currMaxString) - digitsWeCanReplace
			replaceIdx := FindIndexSmallerThanDigit(curr, currMaxString, startIdx)

			if replaceIdx > -1 {
				newMax := currMaxString[:replaceIdx]
				newMax += string(char)
				currMaxString = newMax
				replaced = true
			}

		}

		if len(currMaxString) < targetLength && !replaced {
			currMaxString += string(char)
		}

		fmt.Printf("str=%s, len=%d\n", currMaxString, len(currMaxString))
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

func FindIndexSmallerThanDigit(char int, bank string, start int) int {
	for i := start; i < len(bank); i++ {
		if charToInt(bank[i]) < char {
			return i
		}
	}
	return -1
}
