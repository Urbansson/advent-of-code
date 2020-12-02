package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	input := aoc.ReadInput("./input.txt")
	lines := aoc.ExtractLines(input)
	count := 0
	for _, l := range lines {
		if validate(parse(l)) {
			count++
		}
	}
	fmt.Println("Valid passwords", count)
}

func parse(line string) (int, int, string, string) {
	x := strings.Split(line, ":")
	l := strings.Split(x[0], " ")
	minMax := strings.Split(l[0], "-")
	min := aoc.Atoi(minMax[0])
	max := aoc.Atoi(minMax[1])
	return min, max, l[1], x[1]
}

func validate(min, max int, letter, password string) bool {
	c := strings.Count(password, letter)

	return c >= min && c <= max
}
