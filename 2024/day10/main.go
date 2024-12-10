package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

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

	worldMap.Print(func(v int) string {
		return fmt.Sprintf("%d", v)
	})

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
	// go up
	sum += explore(g, aoc.XY{
		X: s.X,
		Y: s.Y + 1,
	}, i, t)
	// go down
	sum += explore(g, aoc.XY{
		X: s.X,
		Y: s.Y - 1,
	}, i, t)
	// go left
	sum += explore(g, aoc.XY{
		X: s.X - 1,
		Y: s.Y,
	}, i, t)
	// go right
	sum += explore(g, aoc.XY{
		X: s.X + 1,
		Y: s.Y,
	}, i, t)

	return sum
}
