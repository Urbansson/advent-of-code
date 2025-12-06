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

	// Find longest row
	max := 0
	for _, v := range lines {
		max = aoc.Max(max, len(v))
	}

	// Make sure all rows are the same length by padding spaces to the end
	for i, v := range lines {
		if len(v) != max {
			padding := max - len(v)
			lines[i] = v + strings.Repeat(" ", padding)
		}
	}

	n := 0
	currentSum := 0
	totalSum := 0
	for i := 0; i < max; i++ {
		op := ops[n]
		rn := ""
		for _, v := range lines[:len(lines)-1] {
			rn += string(v[i])
		}
		v, err := strconv.Atoi(strings.Trim(rn, " "))
		// If we fail to parse its the devider of an string with only spaces.
		// Reset and add to totalSum
		if err != nil {
			n++
			totalSum += currentSum
			currentSum = 0
			continue
		}
		if op == "*" {
			if currentSum == 0 {
				currentSum++
			}
			currentSum *= v
		} else {
			currentSum += v
		}
	}
	fmt.Println(totalSum + currentSum)
}
