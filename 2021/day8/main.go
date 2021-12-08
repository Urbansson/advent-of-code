package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	res := 0
	for _, line := range lines {
		p := strings.Split(line, " | ")

		outputs := strings.Split(p[1], " ")
		for _, v := range outputs {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				res++
			}
		}
	}
	fmt.Println(res)
}
