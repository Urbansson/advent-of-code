package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	dail := Dail{position: 50, passes: 0}
	for _, l := range lines {
		op := l[0:1]
		v := aoc.Atoi(l[1:])
		if op == "L" {
			for i := 0; i < v; i++ {
				dail.RotateLeft()
			}
		}
		if op == "R" {
			for i := 0; i < v; i++ {
				dail.RotateRight()
			}
		}
	}
	fmt.Printf("pass: %d \n", dail.passes)
}

type Dail struct {
	position int
	passes   int
}

func (d *Dail) RotateLeft() {
	if d.position == 0 {
		d.position = 99
	} else {
		d.position -= 1
		if d.position == 0 {
			d.passes++
		}
	}
}

func (d *Dail) RotateRight() {
	if d.position == 99 {
		d.position = 0
		d.passes++
	} else {
		d.position += 1
	}
}
