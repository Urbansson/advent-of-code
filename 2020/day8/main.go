package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	for i, l := range lines {
		cache := lines[i]
		if strings.HasPrefix(l, "jmp") {
			lines[i] = strings.Replace(l, "jmp", "nop", 1)
		}
		if strings.HasPrefix(l, "nop") {
			lines[i] = strings.Replace(l, "nop", "jmp", 1)
		}
		if acc, ok := run(lines); ok {
			fmt.Println(acc)
			return
		}
		lines[i] = cache
	}
}

func run(input []string) (int, bool) {
	cache := make([]bool, len(input))
	acc := 0
	for i := 0; i < len(input); {
		d := strings.Split(input[i], " ")
		if cache[i] {
			return acc, false
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
	return acc, true
}
