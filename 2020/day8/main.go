package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	cache := make([]bool, len(lines))
	acc := 0
	for i := 0; i < len(lines); {
		d := strings.Split(lines[i], " ")
		if cache[i] {
			break
		}
		cache[i] = true
		switch d[0] {
		case "nop":
			i++
		case "jmp":
			i += aoc.Atoi(d[1])
		case "acc":
			acc += aoc.Atoi(d[1])
			i++
		}
	}
	fmt.Println(acc)
}
