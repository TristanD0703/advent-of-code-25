package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	var builder strings.Builder

	removed := true
	totalRemoved := 0

	for removed {
		removed = false
		for y, line := range lines {
			for x, roll := range line {
				fmt.Printf("%c", roll)
				if roll == '@' && CanRemove(lines, x, y) {
					totalRemoved++
					removed = true
					roll = '.'
				}

				builder.WriteRune(roll)
			}
			fmt.Printf("\n")
			builder.WriteRune('\n')
		}

		lines = strings.Split(builder.String(), "\n")
		builder.Reset()
	}
	fmt.Printf("Total removed rolls: %d\n", totalRemoved)
}

func CanRemove(lines []string, x int, y int) bool {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		currX := dx + x
		if currX < 0 || currX >= len(lines[0]) {
			continue
		}

		for dy := 1; dy >= -1; dy-- {
			currY := dy + y

			if currY < 0 || currY >= len(lines[0]) || (dy == 0 && dx == 0) {
				continue
			}

			if lines[currY][currX] == '@' {
				count++
			}
		}
	}

	if count < 4 {
		return true
	}

	return false
}
