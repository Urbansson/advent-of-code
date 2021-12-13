package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type pos struct {
	x int
	y int
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	code := make(map[pos]bool)
	offset := 0
	for i, l := range lines {
		if len(l) == 0 {
			offset = i + 1
			break
		}
		c := aoc.ExtractInts(l)
		code[pos{x: c[0], y: c[1]}] = true
	}
	printCode(code)
	for _, l := range lines[offset:] {
		fold := strings.Split(l, "=")

		axis := string(fold[0][len(fold[0])-1])
		i := aoc.Atoi(fold[1])

		ncode := make(map[pos]bool)
		for c := range code {

			if axis == "x" {
				if c.x > i {
					ncode[pos{
						x: int(math.Abs(float64((c.x - i) - i))),
						y: c.y,
					}] = true
				} else {
					ncode[c] = true
				}
			} else if axis == "y" {
				if c.y > i {
					ncode[pos{
						y: int(math.Abs(float64((c.y - i) - i))),
						x: c.x,
					}] = true
				} else {
					ncode[c] = true
				}
			}
		}
		code = ncode
	}
	printCode(code)
}

func printCode(code map[pos]bool) {
	for y := 0; y < 150; y++ {
		for x := 0; x < 150; x++ {
			if _, ok := code[pos{
				x: x,
				y: y,
			}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
