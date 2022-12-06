package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdinRaw()
	fmt.Println(findSOP(data, 4))
	fmt.Println(findSOP(data, 14))
}

func findSOP(input string, l int) int {
	for i := 0; i < len(input); i++ {
		sop := input[i : i+l]
		if !hasRepeated(sop) {
			return i + l
		}
	}
	return -1
}

func hasRepeated(s string) bool {
	seen := map[rune]interface{}{}
	for _, c := range s {
		if _, ok := seen[c]; ok {
			return true
		}
		seen[c] = struct{}{}
	}
	return false
}
