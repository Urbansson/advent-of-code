package main

import (
	"fmt"
	"sort"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	opener := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	closer := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	points := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

	scores := []int{}
	for _, l := range lines {
		stack := []rune{}
		corrupted := false
		for _, r := range l {
			if _, ok := opener[r]; ok {
				stack = append(stack, r)
			} else {
				if stack[len(stack)-1] != closer[r] {
					corrupted = true
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
		fmt.Println(len(stack), corrupted)
		if !corrupted {
			var score int
			for i := len(stack) - 1; i >= 0; i-- {
				score = score*5 + points[stack[i]]
			}
			scores = append(scores, score)

		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
