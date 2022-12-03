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

	fmt.Println(total, findBadge(lines))
}

func findBadge(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i += 3 {
		a, b, c := lines[i], lines[i+1], lines[i+2]
		s1 := map[rune]struct{}{}
		s2 := map[rune]struct{}{}
		for _, char := range a {
			s1[char] = struct{}{}
		}
		for _, char := range b {
			s2[char] = struct{}{}
		}
		for _, char := range c {
			_, ok1 := s1[char]
			_, ok2 := s2[char]
			if ok1 && ok2 {
				total += toPrio(char)
				break
			}
		}
	}

	return total
}

func toPrio(r rune) int {
	score := int(r) - 96
	if score < 0 {
		score = int(r) - 64 + 26
	}
	return score
}
