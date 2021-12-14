package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	template := lines[0]
	rules := map[string]string{}
	for _, l := range lines[2:] {
		pair := strings.Split(l, " -> ")
		rules[pair[0]] = pair[1]
	}
	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pairs[string(template[i])+string(template[i+1])]++
	}
	for i := 0; i < 40; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[string(k[0])+rules[k]] += v
			newPairs[rules[k]+string(k[1])] += v
		}
		pairs = newPairs
	}
	counts := map[rune]int{}
	for k, v := range pairs {
		counts[rune(k[0])] += v
	}
	cs := []int{}
	for _, c := range counts {
		cs = append(cs, c)
	}
	sort.Ints(cs)
	fmt.Println(cs[len(cs)-1] - cs[0] - 1)
}
