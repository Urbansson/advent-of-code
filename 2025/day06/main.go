package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	ops := strings.Fields(lines[len(lines)-1])
	numberLines := lines[:len(lines)-1]

	// Normalize line lengths
	maxWidth := 0
	for _, line := range numberLines {
		maxWidth = aoc.Max(maxWidth, len(line))
	}

	for i := range numberLines {
		numberLines[i] += strings.Repeat(" ", maxWidth-len(numberLines[i]))
	}

	opIndex := 0
	currentSum := 0
	totalSum := 0

	for col := 0; col < maxWidth; col++ {
		// Extract vertical number
		columnStr := ""
		for _, line := range numberLines {
			columnStr += string(line[col])
		}

		value, err := strconv.Atoi(strings.TrimSpace(columnStr))

		if err != nil {
			// Space divider - finalize current calculation
			totalSum += currentSum
			currentSum = 0
			opIndex++
			continue
		}

		if opIndex >= len(ops) {
			break
		}

		switch ops[opIndex] {
		case "*":
			if currentSum == 0 {
				currentSum = value
			} else {
				currentSum *= value
			}
		case "+":
			currentSum += value
		}
	}

	fmt.Println(totalSum + currentSum)
}
