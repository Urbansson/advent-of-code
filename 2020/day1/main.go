package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	s := aoc.ReadInput("./input.txt")
	is := aoc.ExtractInts(s)
	if s1, s2, ok := twoSum(is, 2020); ok {
		fmt.Printf("%d+%d=%d\n", s1, s2, s1+s2)
		fmt.Println(s1 * s2)
	} else {
		fmt.Println("target not found")
	}
}

func twoSum(nums []int, target int) (int, int, bool) {
	m := make(map[int]bool, len(nums))
	for _, k := range nums {
		comp := target - k
		if _, ok := m[comp]; ok {
			return comp, k, true
		}
		m[k] = true
	}
	return 0, 0, false
}
