package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	storage := aoc.Grid[rune]{}

	for y, rows := range lines {
		for x, v := range rows {
			xy := aoc.XY{X: x, Y: y}
			storage[xy] = v
		}
	}

	sum := 0
	pr := 0
	for {
		res := aoc.Grid[int]{}
		for xy, v := range storage {
			if v != '@' {
				continue
			}
			check(storage, xy, res)
		}

		if len(res) == pr {
			break
		}
		for k, v := range res {
			if v < 4 {
				storage[k] = '.'
				sum++
			}
		}
		// Print solution part 1
		if pr == 0 {
			fmt.Println(sum)
		}
		pr = len(res)
	}

	// Print solution part 2
	fmt.Println(sum)
}

func check(g aoc.Grid[rune], p aoc.XY, visited aoc.Grid[int]) {
	sum := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}
			v := g[aoc.XY{X: p.X + x, Y: p.Y + y}]
			if v == '@' {
				sum++
			}
		}
	}
	visited[p] = sum
}
