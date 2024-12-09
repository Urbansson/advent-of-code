package aoc

import "fmt"

type XY struct {
	X int
	Y int
}

func Offset(a, b XY) XY {
	return XY{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func PrintGrid[V int64 | float64 | rune](grid map[XY]V, f func(V) string) {
	if len(grid) == 0 {
		fmt.Println("Grid is empty.")
		return
	}

	// Determine bounds of the grid
	var minX, minY, maxX, maxY int

	for coord := range grid {
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
			val, ok := grid[XY{X: x, Y: y}]
			if ok {
				fmt.Print(f(val))
			} else {
				fmt.Print(" ") // Empty cell
			}
		}
		fmt.Println() // Newline for next row
	}
}
