package main

import (
	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	engine := map[aoc.XY]rune{}

	for x, l := range lines {
		for y, v := range l {
			engine[aoc.XY{
				X: x,
				Y: y,
			}] = v
		}
	}

}
