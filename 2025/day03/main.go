package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	sum := 0
	for _, l := range lines {
		sum += findJoltage(12, aoc.DigitList(l))
	}
	fmt.Println(sum)
}

func findJoltage(n int, s []int) int {
	var val []int
	si := -1
	for i := n; i > 0; i-- {
		var max int
		var maxIndex int
		for j := si + 1; j <= len(s)-i; j++ {
			if s[j] > max {
				max = s[j]
				maxIndex = j
			}
		}
		val = append(val, max)
		si = maxIndex
	}
	return aoc.DigitsToInt(val)
}
