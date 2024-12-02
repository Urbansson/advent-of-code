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
		il := aoc.ExtractInts(l)

		if safeCheck((il)) {
			sum++
		}

	}

	fmt.Println("sum", sum)
}

func safeCheck(i []int) bool {
	valid := true

	// indicates if report should increase or deacrese
	dir := 0

	p := i[0]
	for _, v := range i[1:] {
		if p < v {
			if dir == 0 {
				dir = 1
			}
			if dir != 1 {
				valid = false
				break
			}

			if !safeStep(v - p) {
				valid = false
				break
			}

		} else if p > v {
			if dir == 0 {
				dir = -1
			}
			if dir != -1 {
				valid = false
				break
			}
			if !safeStep(p - v) {
				valid = false
				break
			}

		} else {
			// Not increasing or decreasing == invalid
			valid = false
			break
		}
		p = v
	}
	return valid
}

func safeStep(i int) bool {
	return 0 < i && i <= 3
}
