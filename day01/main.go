package main

import (
	"aoc/utils"
	"fmt"
	"math"
)

func main() {
	input := utils.ReadInput("/Users/theodor/Documents/svt/aoc/day01/input.txt")
	data := utils.IntList(input)
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(input []int) int {
	sum := 0
	for _, v := range input {
		sum += calcFuel(v)
	}
	return sum
}

func part2(input []int) int {
	sum := 0
	for _, v := range input {
		mod := 0
		n := calcFuel(v)
		for n > 0 {
			mod += n
			n = calcFuel(n)
		}
		sum += mod
	}
	return sum
}

func calcFuel(mass int) int {
	return int(math.Max(float64((mass/3)-2), 0))
}
