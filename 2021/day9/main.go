package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	cave := make([][]int, len(lines))
	for i, l := range lines {
		r := aoc.DigitList(l)
		cave[i] = r
	}

	res := 0
	for x := range cave {
		for y := range cave[x] {
			height := cave[x][y]
			var above, left, under, right = 9, 9, 9, 9
			if x-1 >= 0 {
				above = cave[x-1][y]
			}
			if y-1 >= 0 {
				left = cave[x][y-1]
			}
			if x+1 < len(cave) {
				under = cave[x+1][y]
			}
			if y+1 < len(cave[x]) {
				right = cave[x][y+1]
			}
			if height < above && height < left && height < under && height < right {
				fmt.Println(x, y, height)
				res += height + 1
			}
		}
	}
	fmt.Println(res)
}
