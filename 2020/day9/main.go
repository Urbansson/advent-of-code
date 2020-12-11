package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

const preamble = 25

func main() {
	data := aoc.ReadStdin()
	num := aoc.ExtractInts(data)
	for i := preamble; i < len(num); i++ {
		pre := num[i-preamble : i]
		if _, _, ok := twoSum(pre, num[i]); !ok {
			fmt.Println(num[i])

			min, max := aoc.MinMax(findSeq(num, num[i]))
			fmt.Println(min, max, min+max)
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

func findSeq(nums []int, target int) []int {
	for i, v := range nums {
		sum := v
		vals := []int{v}
		for _, k := range nums[i+1:] {
			sum += k
			vals = append(vals, k)
			if sum == target {
				return vals
			}
			if sum > target {
				break
			}
		}
	}
	return nil
}
