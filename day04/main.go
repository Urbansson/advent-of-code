package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("day04/input.txt")
	data := utils.ExtractInts(input)
	count := 0
	for i := data[0]; i < data[1]; i++ {
		if monotonic(i) && nextEqual(i) {
			count++
		}
	}
	fmt.Println(count)
}

func monotonic(i int) bool {
	digits := Digits(i)
	for k := 0; k < len(digits)-1; k++ {
		if digits[k] > digits[k+1] {
			return false
		}
	}
	return true
}

func nextEqual(i int) bool {
	digits := Digits(i)
	for k := 0; k < len(digits)-1; k++ {
		if digits[k] == digits[k+1] {
			return true
		}
	}
	return false
}

func Digits(i int) []int {
	var digits []int
	for i != 0 {
		digits = append([]int{i % 10}, digits...)
		i /= 10
	}
	return digits
}
