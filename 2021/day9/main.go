package main

import (
	"fmt"
	"sort"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type pos struct {
	x int
	y int
}

var adjacent = []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	cave := map[pos]int{}
	for x, l := range lines {
		for y, h := range aoc.DigitList(l) {
			cave[pos{x: x, y: y}] = h
		}
	}
	res := 0
	var lowpoints []pos
	for p, height := range cave {
		low := 9
		for _, a := range adjacent {
			if ah, ok := cave[pos{
				x: p.x + a.x,
				y: p.y + a.y,
			}]; ok {
				if ah < low {
					low = ah
				}
			}
		}
		if height < low {
			res += height + 1
			lowpoints = append(lowpoints, p)
		}
	}
	fmt.Println(res)
	var sums []int
	for _, lowpoint := range lowpoints {
		s := larger(cave, lowpoint, lowpoint, make(map[pos]bool))
		sums = append(sums, s)
	}
	sort.Ints(sums)
	fmt.Println(sums[len(sums)-3] * sums[len(sums)-2] * sums[len(sums)-1])

}

func larger(grid map[pos]int, start pos, root pos, seen map[pos]bool) (sum int) {
	sum += 1
	for _, a := range adjacent {
		newP := pos{
			x: start.x + a.x,
			y: start.y + a.y,
		}
		if _, ok := seen[newP]; ok {
			continue
		}
		if vv, ok := grid[newP]; ok {
			if vv > grid[root] && vv != 9 {
				seen[newP] = true
				sum += larger(grid, newP, root, seen)
			}
		}
	}
	return
}
