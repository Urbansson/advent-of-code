package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	opener := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	closer := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

	res := 0
	for _, l := range lines {
		stack := []rune{}
		for _, r := range l {
			if _, ok := opener[r]; ok {
				stack = append(stack, r)
			} else {
				if stack[len(stack)-1] != closer[r] {
					res += points[r]
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	fmt.Println(res)
}
