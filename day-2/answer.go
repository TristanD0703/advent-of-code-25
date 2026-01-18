package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)
	ranges := strings.SplitSeq(file, ",")
	invalidCount := 0

	for idRange := range ranges {
		ret := strings.Split(idRange, "-")

		start, err := strconv.Atoi(ret[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(ret[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			curr := strconv.Itoa(i)

			if !CheckProductID(curr) {
				invalidCount += i
				fmt.Printf("Found invalid ID: %s\n", curr)
			}
		}
	}

	fmt.Printf("Invalid IDs: %d\n", invalidCount)
}

// Checks if a product id is invalid according to rules. Returns false for invalid.
func CheckProductID(id string) bool {

	// Sliding window approach - Increase the repeated sequence substring length until it repeats over the whole string
	for seqLength := 1; seqLength <= len(id)/2; seqLength++ {
		currSeq := id[:seqLength]
		curr := seqLength
		repeat := true
		count := 1

		// Compare beginning sequence to all other non-overlapping sequences
		for repeat && curr+seqLength <= len(id) {
			subseqToComp := id[curr : curr+seqLength]

			if strings.EqualFold(currSeq, subseqToComp) {
				count++
			} else {
				repeat = false
			}
			curr += seqLength
		}

		// Check if substring repeats over the whole string
		if float32(count) == float32(len(id))/float32(seqLength) {
			return false
		}
	}

	// If we make it here, there are no repeating subsequences
	return true
}
