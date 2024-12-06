package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type guard struct {
	direction rune // < > ^ v
	pos       aoc.XY

	uniqueSpots int
	steps       int
	visited     map[aoc.XY]rune
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	g := map[aoc.XY]rune{}
	bs := []aoc.XY{}

	var gp guard
	for y, l := range lines {
		for x, v := range l {
			p := aoc.XY{
				X: x,
				Y: y,
			}
			g[p] = v

			if v == '.' {
				bs = append(bs, p)
			}

			if v == '<' || v == '>' || v == '^' || v == 'v' {
				gp = guard{
					pos:         p,
					direction:   v,
					uniqueSpots: 1,
					visited:     make(map[aoc.XY]rune),
				}
				g[p] = 'X'
			}
		}
	}

	PrettyPrintGrid(g)

	validBlocks := 0
	for _, p := range bs {
		ng := make(map[aoc.XY]rune)
		for k, v := range g {
			ng[k] = v
		}
		cpg := gp.clone()

		// Add the new block
		ng[p] = '#'

		for {
			bounds, cycle := cpg.walk(ng)
			// To many steps, stuck in loop
			if cycle {
				validBlocks++
				break
			}
			if !bounds {
				break
			}
		}

	}
	fmt.Println("blocks", validBlocks)
}

func (g *guard) clone() guard {
	return guard{
		pos:         g.pos,
		direction:   g.direction,
		uniqueSpots: 1,
		steps:       0,
		visited:     make(map[aoc.XY]rune),
	}
}

func (g *guard) changeDirection() {
	switch g.direction {
	case '^':
		g.direction = '>'
	case 'v':
		g.direction = '<'
	case '<':
		g.direction = '^'
	case '>':
		g.direction = 'v'
	}
}

func (g *guard) walk(grid map[aoc.XY]rune) (bool, bool) {
	var np aoc.XY

	switch g.direction {
	case '^':
		np = aoc.XY{
			X: g.pos.X,
			Y: g.pos.Y - 1,
		}
	case 'v':
		np = aoc.XY{
			X: g.pos.X,
			Y: g.pos.Y + 1,
		}
	case '<':
		np = aoc.XY{
			X: g.pos.X - 1,
			Y: g.pos.Y,
		}
	case '>':
		np = aoc.XY{
			X: g.pos.X + 1,
			Y: g.pos.Y,
		}
	}

	if dir, ok := g.visited[np]; ok && dir == g.direction {
		return false, true
	}

	nextStep, ok := grid[np]
	// Outside of bounds, indicate we could not walk
	if !ok {
		return false, false
	}

	// way blocked
	if nextStep == '#' {
		g.changeDirection()
		return true, false
	}

	// Mark as visited
	grid[np] = 'X'

	if nextStep == '.' {
		g.uniqueSpots += 1
	}

	g.visited[np] = g.direction
	g.steps += 1
	// Move
	g.pos = np

	return true, false
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
