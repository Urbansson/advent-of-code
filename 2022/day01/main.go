package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	max := 0
	temp := 0
	for _, l := range lines {
		if l == "" {
			max = aoc.Max(max, temp)
			temp = 0
		} else {
			temp += aoc.Atoi(l)
		}
	}
	fmt.Println(max)
}
