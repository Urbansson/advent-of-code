package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	data := utils.ReadInput("day02/input.txt")
	prog := utils.ExtractInts(data)
	// Part1
	res := execute(prog, 12, 2)
	fmt.Println(res)

	// Part2
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			res := execute(prog, n, v)
			if res == 19690720 {
				fmt.Printf("%d%d", n, v)
				return
			}
		}
	}
}

func execute(p []int, noun, verb int) int {
	p = append([]int(nil), p...)
	p[1] = noun
	p[2] = verb
	for i := 0; i < len(p); i += 4 {
		op, x, y, z := p[i], p[i+1], p[i+2], p[i+3]
		switch op {
		case 1:
			p[z] = p[x] + p[y]
		case 2:
			p[z] = p[x] * p[y]
		case 99:
			return p[0]
		default:
			panic("invalid op code")
		}
	}
	panic("syntax error")
}
