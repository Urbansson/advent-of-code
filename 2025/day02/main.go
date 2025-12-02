package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	ids := strings.Split(data, ",")
	sum := 0
	for _, id := range ids {
		parts := strings.Split(id, "-")
		p1 := aoc.Atoi(parts[0])
		p2 := aoc.Atoi(parts[1])
		for i := p1; i <= p2; i++ {
			v := strconv.Itoa(i)
			for j := 1; j <= len(v)/2; j++ {
				if strings.Repeat(v[0:j], len(v)/j) == v {
					sum += i
					break
				}
			}
		}
	}
	fmt.Println(sum)
}
