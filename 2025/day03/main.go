package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	sum := 0
	for _, l := range lines {
		bs := aoc.DigitList(l)
		max := 0
		for i, j := range bs {
			for _, k := range bs[i+1:] {
				c := aoc.DigitsToInt([]int{j, k})
				max = aoc.Max(max, c)
			}
		}
		sum += max
	}
	fmt.Println(sum)
}
