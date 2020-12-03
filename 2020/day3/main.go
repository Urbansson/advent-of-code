package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	input := aoc.ReadInput("./input.txt")
	lines := aoc.ExtractLines(input)

	fmt.Println(navigate(lines, 3, 1))

	out := navigate(lines, 1, 1) *
		navigate(lines, 3, 1) *
		navigate(lines, 5, 1) *
		navigate(lines, 7, 1) *
		navigate(lines, 1, 2)

	fmt.Println(out)
}

func navigate(input []string, r, d int) int {
	x := 0
	c := 0
	for i := d; i < len(input); i += d {
		x += r
		if x >= len(input[i]) {
			x = x - len(input[i])
		}
		if string(input[i][x]) == "#" {
			c++
		}
	}
	return c
}
