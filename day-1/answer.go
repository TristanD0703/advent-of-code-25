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
	defer file.Close()

	reader := bufio.NewScanner(file)
	var previousPointer int64 = 0
	var currPointer int64 = 50
	var password int64 = 0

	for reader.Scan() {
		line := reader.Text()

		// Parse the number to move pointer
		num, err := strconv.ParseInt(line[1:], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		if num == 0 {
			continue
		}

		direction := line[0]

		// Going left - subtract from curr position
		if direction == 'L' {
			num *= -1
		} else if direction != 'R' {
			log.Fatal("Invalid rotation character")
		}

		previousPointer = currPointer
		currPointer += num
		rotations := abs(currPointer / 100) // Calculate how many overflows

		// For some reason -1 % 100 == -1??
		currPointer %= 100

		// Handle pointer underflow due to behavior of % operator on - nums
		if currPointer < 0 {
			// ignore if we left from 0 last time
			if previousPointer != 0 {
				rotations += 1
			}

			// Make it positive again
			currPointer += 100

			// Check if we hit 0 exactly. Ignore if rotating right as checking how many overflows accounts for this
		} else if currPointer == 0 && direction == 'L' {
			rotations += 1
		}

		fmt.Printf("in=%s, currPointer=%d, rotations=%d, ", line, currPointer, rotations)

		password += rotations
		fmt.Printf("password=%d\n", password)
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The password is %d\n", password)
}

func abs(num int64) int64 {
	if num < 0 {
		num *= -1
	}
	return num
}
