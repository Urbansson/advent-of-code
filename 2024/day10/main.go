package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

var directions = []aoc.XY{
	{X: 1, Y: 0},  // Right
	{X: -1, Y: 0}, // Left
	{X: 0, Y: 1},  // Down
	{X: 0, Y: -1}, // Up
}

func main() {
	lines := aoc.ExtractLines(aoc.ReadStdin())

	heads := []aoc.XY{}
	worldMap := aoc.Grid[int]{}
	for y, rows := range lines {
		for x, v := range rows {
			xy := aoc.XY{X: x, Y: y}

			inc := aoc.Atoi(string(v))
			if inc == 0 {
				heads = append(heads, xy)
			}
			worldMap[xy] = inc
		}
	}

	sum1 := 0
	sum2 := 0
	for _, h := range heads {
		t := make(map[aoc.XY]int)
		sum1 += explore(worldMap, h, -1, t)
		sum2 += len(t)
	}

	fmt.Println("---")
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func explore(g aoc.Grid[int], s aoc.XY, prev int, t map[aoc.XY]int) int {
	i, ok := g[s]
	if !ok {
		return 0
	}

	if i != prev+1 {
		return 0
	}

	// Reached the end of a trail
	if i == 9 {
		t[s] += 1
		return 1
	}

	sum := 0
	for _, d := range directions {
		sum += explore(g, aoc.XY{
			X: s.X + d.X,
			Y: s.Y + d.Y,
		}, i, t)
	}

	return sum
}
