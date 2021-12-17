package main

import (
	_ "embed"
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	var x1, x2, y1, y2 int
	fmt.Sscanf(data, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	calc := func(v aoc.XY) (int, bool) {
		top := 0
		pos := aoc.XY{}
		for {
			pos.X += v.X
			pos.Y += v.Y
			if pos.Y > top {
				top = pos.Y
			}
			if pos.X >= x1 && pos.X <= x2 && pos.Y >= y1 && pos.Y <= y2 {
				return top, true
			}
			if v.X == 0 && v.Y < y1 {
				return 0, false
			}
			if v.X > 0 {
				v.X--
			}
			if v.X < 0 {
				v.X++
			}
			v.Y--
		}
	}
	var top int
	b := 500
	r := make(map[aoc.XY]bool)
	for x := b * -1; x < b; x++ {
		for y := b * -1; y < b; y++ {
			v := aoc.XY{X: x, Y: y}
			if m, ok := calc(v); ok {
				if m > top {
					top = m
				}
				r[v] = true
			}
		}
	}
	fmt.Println(top, len(r))
}
