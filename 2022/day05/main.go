package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type Stack[C any] []C

func (s *Stack[C]) Push(v C) {
	*s = append(*s, v)
}

// PushN pushes n elements from the stack to the stack
func (s *Stack[C]) PushN(v []C) {
	*s = append(*s, v...)
}

func (s *Stack[C]) Pop() C {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[C]) PopN(n int) []C {
	if len(*s) < n {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return v
}

func (s *Stack[C]) Peek() C {
	return (*s)[len(*s)-1]
}

func main() {
	data := aoc.ReadStdinRaw()
	split := strings.Split(data, "\n\n")
	c := aoc.ExtractLinesRaw(split[0])
	rs := make([]Stack[rune], len(c[0]))
	for i := len(c) - 1; i >= 0; i-- {
		for i, r := range c[i] {
			if unicode.IsLetter(r) {
				rs[i].Push(r)
			}
		}
	}

	stacks := []Stack[rune]{}
	for _, s := range rs {
		if len(s) != 0 {
			stacks = append(stacks, s)
		}
	}

	for _, l := range aoc.ExtractLines(split[1]) {
		var amount, from, to int
		fmt.Sscanf(l, "move %d from %d to %d", &amount, &from, &to)
		stacks[to-1].PushN(stacks[from-1].PopN(amount))
	}

	for _, s := range stacks {
		fmt.Print(string(s.Peek()))
	}
	fmt.Println()
}
