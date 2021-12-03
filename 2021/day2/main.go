package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/2021/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	x, y, aim := 0, 0, 0
	for _, l := range lines {
		var op string
		var v int
		fmt.Sscanf(l, "%s %d", &op, &v)
		switch op {
		case "forward":
			x += v
			y += (aim * v)
		case "down":
			aim += v
		case "up":
			aim -= v
		}
	}
	fmt.Println(x * y)
}
