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
		if validateToboggan(parse(l)) {
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
	return min, max, strings.TrimSpace(l[1]), strings.TrimSpace(x[1])
}

func validate(min, max int, letter, password string) bool {
	c := strings.Count(password, letter)

	return c >= min && c <= max
}

func validateToboggan(first, second int, letter, password string) bool {
	l1 := password[first-1]
	l2 := password[second-1]
	if l1 != l2 && (l1 == letter[0] || l2 == letter[0]) {
		return true
	}
	return false
}
