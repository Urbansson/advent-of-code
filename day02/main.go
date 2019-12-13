package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	data := utils.ReadInput("day02/input.txt")
	prog := utils.ExtractInts(data)
	res := execute(prog)
	fmt.Println(res)
}

func execute(p []int) int {
	p = append([]int(nil), p...)
	p[1] = 12
	p[2] = 2
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
