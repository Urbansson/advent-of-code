package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/2021/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	var h, v int
	for i, l := range lines {
		fmt.Println(i, l)
		f := strings.Split(l, " ")
		d := aoc.Atoi(f[1])
		switch f[0] {
		case "forward":
			v = v + d
		case "down":
			h = h + d
		case "up":
			h = h - d
		}
	}
	fmt.Println(h * v)
}
