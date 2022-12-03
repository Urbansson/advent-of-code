package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	total := 0
	for _, r := range lines {
		set := map[rune]struct{}{}
		first, last := r[:len(r)/2], r[len(r)/2:]

		for _, char := range first {
			set[char] = struct{}{}
		}
		for _, char := range last {
			if _, ok := set[char]; ok {
				total += toPrio(char)
				break
			}
		}
	}
	fmt.Println(total)
}

func toPrio(r rune) int {
	score := int(r) - 96
	if score < 0 {
		score = int(r) - 64 + 26
	}
	return score
}
