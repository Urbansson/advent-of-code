package main

import (
	"fmt"
	"math"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	grid := map[aoc.XY]rune{}

	for y, l := range lines {
		for x, v := range l {
			grid[aoc.XY{
				X: x,
				Y: y,
			}] = v
		}
	}

	PrettyPrintGrid(grid)

	fmt.Println("---")
	fmt.Println(Solve(grid))
	fmt.Println(Solve2(grid))
}

func Solve2(grid map[aoc.XY]rune) int {
	sum := 0

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
			masCount := 0
			if grid[aoc.XY{X: x, Y: y}] == 'A' {
				if grid[aoc.XY{X: x - 1, Y: y - 1}] == 'M' && grid[aoc.XY{X: x + 1, Y: y + 1}] == 'S' {
					masCount++
				}

				if grid[aoc.XY{X: x - 1, Y: y - 1}] == 'S' && grid[aoc.XY{X: x + 1, Y: y + 1}] == 'M' {
					masCount++
				}

				if grid[aoc.XY{X: x - 1, Y: y + 1}] == 'M' && grid[aoc.XY{X: x + 1, Y: y - 1}] == 'S' {
					masCount++
				}

				if grid[aoc.XY{X: x - 1, Y: y + 1}] == 'S' && grid[aoc.XY{X: x + 1, Y: y - 1}] == 'M' {
					masCount++
				}
			}
			if masCount == 2 {
				sum++
			}
		}
	}

	return sum
}

func Solve(grid map[aoc.XY]rune) int {
	sum := 0

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

	directions := []struct {
		dx, dy int
	}{
		{1, 0},   // Right
		{1, 1},   // Diagonal down-right
		{1, -1},  // Diagonal up-right
		{-1, 0},  // Left
		{-1, 1},  // Diagonal down-left
		{-1, -1}, // Diagonal up-left
		{0, 1},   // Down
		{0, -1},  // Up
	}

	// Print the grid
	for y := minY; y <= maxY; y++ { // Start from maxY to print top-down
		for x := minX; x <= maxX; x++ {
			if grid[aoc.XY{X: x, Y: y}] == 'X' {
				for _, dir := range directions {
					// Compute the endpoint of the substring
					to := aoc.XY{X: x + 3*dir.dx, Y: y + 3*dir.dy}
					if subString(grid, aoc.XY{X: x, Y: y}, to) == "XMAS" {
						sum++
					}
				}
			}
		}
	}

	return sum
}

func subString(grid map[aoc.XY]rune, from aoc.XY, to aoc.XY) string {
	var result string
	dx := int(math.Abs(float64(to.X - from.X)))
	dy := int(math.Abs(float64(to.Y - from.Y)))
	sx := 1
	if from.X > to.X {
		sx = -1
	}
	sy := 1
	if from.Y > to.Y {
		sy = -1
	}

	err := dx - dy
	current := from

	for {
		// Append the current cell's value or a space if empty
		if val, ok := grid[current]; ok {
			result += string(val)
		} else {
			result += " "
		}

		// Check if we've reached the target
		if current == to {
			break
		}

		// Update coordinates using Bresenham's algorithm
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			current.X += sx
		}
		if e2 < dx {
			err += dx
			current.Y += sy
		}
	}
	return result
}

func PrettyPrintGrid(grid map[aoc.XY]rune) {
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
			val, ok := grid[aoc.XY{X: x, Y: y}]
			if ok {
				fmt.Print(string(val))
			} else {
				fmt.Print(" ") // Empty cell
			}
		}
		fmt.Println() // Newline for next row
	}
}
