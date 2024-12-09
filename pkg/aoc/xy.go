package aoc

import "fmt"

type XY struct {
	X int
	Y int
}

type Grid[T any] map[XY]T

func Offset(a, b XY) XY {
	return XY{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (g Grid[T]) Print(f func(T) string) {
	if len(g) == 0 {
		fmt.Println("Grid is empty.")
		return
	}

	// Determine bounds of the grid
	var minX, minY, maxX, maxY int

	for coord := range g {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	// Print the grid
	for y := minY; y <= maxY; y++ { // Start from maxY to print top-down
		for x := minX; x <= maxX; x++ {
			val, ok := g[XY{X: x, Y: y}]
			if ok {
				fmt.Print(f(val))
			} else {
				fmt.Print(" ") // Empty cell
			}
		}
		fmt.Println() // Newline for next row
	}
}
