package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0
	for _, l := range lines {
		il := aoc.ExtractInts(l)

		if isSafeRemove((il)) {
			sum++
		}
	}
	fmt.Println("sum", sum)
}

func isSafeRemove(report []int) bool {
	if safeCheck(report) {
		return true
	}

	for i := range report {
		portion := slices.Delete(slices.Clone(report), i, i+1)
		if safeCheck(portion) {
			return true
		}
	}
	return false
}

func safeCheck(input []int) bool {
	isIncreasing := false
	isDecreasing := false

	prev := input[0]

	for _, curr := range input[1:] {
		slc := safeLavelChange(int(math.Abs(float64(prev - curr))))
		if curr > prev {
			isIncreasing = true
		}
		if curr < prev {
			isDecreasing = true
		}
		prev = curr
		// We cant be both increaseing and decrease
		if !slc || (isIncreasing && isDecreasing) {
			return false
		}
	}
	return true
}

func safeLavelChange(i int) bool {
	return 0 < i && i <= 3
}
