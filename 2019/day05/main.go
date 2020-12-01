package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	data := utils.ReadInput("day05/input.txt")
	prog := utils.ExtractInts(data)
	// Part1
	fmt.Println(execute(prog, 1))
	fmt.Println(execute(prog, 5))
}

type Operation int

const (
	Add Operation = iota + 1
	Multiply
	Read
	Write
	JumpTrue
	JumpFalse
	LessThan
	Equals
	Exit = 99
)

type Mode int

const (
	Position Mode = iota
	Immediate
)

func execute(p []int, input int) (output []int) {
	p = append([]int(nil), p...)
	for i := 0; i < len(p); {
		op, flags := parseOpcode(p[i])
		read := func(k int) int {
			mode := Position
			if len(flags) >= k {
				mode = flags[len(flags)-k]
			}
			switch mode {
			case Position:
				return p[p[i+k]]
			case Immediate:
				return p[i+k]
			default:
				panic("invalid mode")
			}
		}
		switch op {
		case Add:
			p[p[i+3]] = read(1) + read(2)
			i += 4
		case Multiply:
			p[p[i+3]] = read(1) * read(2)
			i += 4
		case Read:
			p[p[i+1]] = input
			i += 2
		case Write:
			output = append(output, read(1))
			i += 2
		case JumpTrue:
			if read(1) != 0 {
				i = read(2)
			} else {
				i += 3
			}
		case JumpFalse:
			if read(1) == 0 {
				i = read(2)
			} else {
				i += 3
			}
		case LessThan:
			if read(1) < read(2) {
				p[p[i+3]] = 1
			} else {
				p[p[i+3]] = 0
			}
			i += 4
		case Equals:
			if read(1) == read(2) {
				p[p[i+3]] = 1
			} else {
				p[p[i+3]] = 0
			}
			i += 4
		case Exit:
			return
		default:
			panic("invalid op code")
		}
	}
	panic("syntax error")
}

func parseOpcode(c int) (Operation, []Mode) {
	op := Operation(c % 100)
	var flags []Mode
	for _, m := range utils.Digits(c / 100) {
		flags = append(flags, Mode(m))
	}
	return op, flags
}
