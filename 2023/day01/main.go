package main

import (
	"fmt"
	"unicode"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0
	for _, l := range lines {
		first := 0
		last := 0
		for _, c := range l {
			if ok := unicode.IsDigit(c); ok {
				if first == 0 {
					first = int(c-'0') * 10
				}
				last = int(c - '0')
			}
		}
		sum += first + last

	}

	fmt.Println(sum)
}
