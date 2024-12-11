package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type Entry struct {
	rock  int
	depth int
}

func main() {
	stones := aoc.ExtractInts(aoc.ReadStdin())

	dp := map[Entry]int{}

	var sum int
	for _, v := range stones {
		sum += (applyRule(v, 75, dp))
	}
	fmt.Println("---")
	fmt.Println(sum)
}

func applyRule(rock int, depth int, cache map[Entry]int) int {
	if res, ok := cache[Entry{rock, depth}]; ok {
		return res
	}
	if rock == -1 {
		return 0
	}
	if depth == 0 {
		return 1
	}
	left, right := -1, -1
	if rock == 0 {
		left = 1
	} else if d := aoc.Digits(rock); len(d)%2 == 0 {
		h := len(d) / 2
		left = aoc.DigitsToInt(d[:h])
		right = aoc.DigitsToInt(d[h:])
	} else {
		left = rock * 2024
	}
	cache[Entry{rock, depth}] = applyRule(left, depth-1, cache) + applyRule(right, depth-1, cache)
	return cache[Entry{rock, depth}]
}
