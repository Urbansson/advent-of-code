package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type pos struct {
	x int
	y int
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	template := lines[0]
	pairs := make(map[string]string)
	for _, l := range lines[2:] {
		pair := strings.Split(l, " -> ")
		pairs[pair[0]] = pair[1]
	}
	for i := 0; i < 40; i++ {
		res := template
		offset := 0
		for k := 0; k < len(template)-1; k++ {
			pair := template[k : k+2]
			if pol, ok := pairs[pair]; ok {
				res = res[:k+offset+1] + pol + res[k+offset+1:]
				offset++
			}
		}
		template = res
	}
	mc := make(map[rune]int)
	for _, e := range template {
		mc[e]++
	}
	mcs := []int{}
	for _, c := range mc {
		mcs = append(mcs, c)
	}
	sort.Ints(mcs)
	fmt.Println(mcs[len(mcs)-1] - mcs[0])
}
