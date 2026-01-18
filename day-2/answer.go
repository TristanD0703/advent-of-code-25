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

	CheckProductID("2121212118")
	fmt.Printf("Invalid IDs: %d\n", invalidCount)
}

// Checks if a product id is invalid according to rules. Returns false for invalid.
func CheckProductID(id string) bool {
	// quasai hash table int -> bool
	// var seen [10]bool
	// asciiBeginInts := 48

	// curr := 0
	// for curr < len(id) {
	// 	// Get integer from ascii
	// 	integer := id[curr] % byte(asciiBeginInts)
	// 	if seen[integer] {
	// 		break
	// 	}

	// 	seen[integer] = true
	// 	curr++
	// }

	for seqLength := 1; seqLength <= len(id)/2; seqLength++ {
		currSeq := id[:seqLength]
		curr := seqLength
		repeat := true
		count := 1
		for repeat && curr+seqLength <= len(id) {
			subseqToComp := id[curr : curr+seqLength]

			if strings.EqualFold(currSeq, subseqToComp) {
				count++
			} else {
				repeat = false
			}
			curr += seqLength
		}

		if float32(count) == float32(len(id))/float32(seqLength) {
			return false
		}
	}

	return true
}
