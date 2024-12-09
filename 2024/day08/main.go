package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0

	freq := map[rune][]aoc.XY{}
	worldMap := map[aoc.XY]rune{}
	for y, rows := range lines {
		for x, v := range rows {
			xy := aoc.XY{X: x, Y: y}
			if v != '.' {
				freq[v] = append(freq[v], xy)
			}
			worldMap[xy] = v
		}
	}

	antinodes := map[aoc.XY]bool{}
	for _, a := range freq {
		a := findAntinodes(worldMap, a)
		for k, v := range a {
			if v {
				antinodes[k] = true
			}
		}
	}

	aoc.PrintGrid(worldMap, func(v rune) string {
		return string(v)
	})

	fmt.Println("---")
	fmt.Println(sum, len(antinodes))
}

func findAntinodes(wm map[aoc.XY]rune, antennas []aoc.XY) map[aoc.XY]bool {
	antinodes := make(map[aoc.XY]bool, 0)
	for _, a := range antennas {
		for _, v := range antennas {
			if a == v {
				continue
			}
			offset := aoc.Offset(a, v)
			antinode := aoc.XY{
				X: a.X + offset.X,
				Y: a.Y + offset.Y,
			}
			if v, ok := wm[antinode]; ok {
				if v == '.' {
					wm[antinode] = '#'
				}
				antinodes[antinode] = true
			}
		}
	}
	return antinodes
}
