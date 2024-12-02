package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type Tree struct {
	visable bool
	height  int
	aoc.XY
}

func main() {
	data := aoc.ReadStdinRaw()
	lines := aoc.ExtractLines(data)

	grid := make(map[int]map[int]Tree)

	for x, line := range lines {
		//grid[] = make(map[int]Tree)
		for y := range line {
			grid[y][x] = Tree{
				visable: true,
				height:  aoc.Atoi(string(line[y])),
				XY:      aoc.XY{X: y, Y: x},
			}
		}
	}

	//heihgestT := map[int]int{}
	//heighestB := map[int]int{}
	for _, x := range grid {
		//heighestL := 0
		//From the left find non visable trees.
		for _, y := range x {
			fmt.Print(y.height)
		}
		fmt.Println()

		/*if y.height < heighestL {
				y.visable = false
			}
			heighestL = aoc.Max(heighestL, y.height)
		}
		heighestR := 0
		//From the right find non visable trees.
		for i := len(x); i >= 0; i-- {
			if x[i].height < heighestR {
				y := x[i]
				y.visable = false
			}
			heighestR = aoc.Max(heighestR, x[i].height)
		}
		fmt.Println()
		//		heihgestT[i] = x[0].height
		//Find the highest tree in each column.

		/*
			if heihgestT[x.X] < heighestL {
				heihgestT[x[0].X] = heighestL
			}*/

	}

}
