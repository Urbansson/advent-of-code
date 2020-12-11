package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	num := aoc.ExtractInts(data)
	for i := 25; i < len(num); i++ {
		pre := num[i-25 : i]
		if _, _, ok := twoSum(pre, num[i]); !ok {
			fmt.Println(num[i])
			return
		}
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
