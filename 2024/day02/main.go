package main

import (
	"fmt"
	"math"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0
	for _, l := range lines {
		il := aoc.ExtractInts(l)

		if safeCheck((il)) {
			//fmt.Println("safe", il)
			sum++

		}
		fmt.Println("---")
	}

	fmt.Println("sum", sum)
}

func safeCheck(i []int) bool {
	fmt.Println("testing", i)
	dir := i[0] - i[len(i)-1]
	p := i[0]
	for _, c := range i[1:] {
		fmt.Println("check if safe level", p, c)
		fmt.Println("check step size", int(math.Abs(float64(p-c))), safeLavelChange(int(math.Abs(float64(p-c)))))
		if !safeLavelChange(int(math.Abs(float64(p - c)))) {
			fmt.Println("[Not Valid step size]")
			return false
		}
		// We expect the value to increase
		if dir < 0 {
			fmt.Println("value should increase", p, ">", c, p > c)
			if p > c {
				fmt.Println("[Wrong direction, should increase]")
				return false
			}
		} else {
			fmt.Println("value should decrease", p, "<", c, p < c)
			if p < c {
				fmt.Println("[Wrong direction, should decrease]")
				return false
			}
		}
		p = c
	}
	fmt.Println("Valid report")
	return true
}

func safeLavelChange(i int) bool {
	return 0 < i && i <= 3
}
