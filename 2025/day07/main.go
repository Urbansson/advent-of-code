package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	var start aoc.XY
	grid := aoc.Grid[rune]{}
	for y, line := range lines {
		for x, v := range line {
			if v == 'S' {
				start = aoc.XY{X: x, Y: y}
			}
			grid[aoc.XY{X: x, Y: y}] = v
		}
	}
	visited := make(map[aoc.XY]bool)
	fmt.Println(traverse(aoc.XY{X: start.X, Y: start.Y + 1}, grid, visited))
}

func traverse(s aoc.XY, g aoc.Grid[rune], visited map[aoc.XY]bool) int {
	v := g[s]

	if visited[s] {
		return 0
	}

	if v == '.' {
		visited[s] = true
		return traverse(aoc.XY{X: s.X, Y: s.Y + 1}, g, visited)
	}

	if v == '^' {
		left := traverse(aoc.XY{X: s.X - 1, Y: s.Y}, g, visited)
		right := traverse(aoc.XY{X: s.X + 1, Y: s.Y}, g, visited)
		return left + right + 1
	}

	return 0
}
