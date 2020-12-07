package main

import (
	"fmt"
	"regexp"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	containedin := make(map[string]map[string]bool, 0)
	r, _ := regexp.Compile("(.+?) bags contain")
	ir, _ := regexp.Compile("(\\d+) (.+?) bags?[,.]")
	for _, line := range aoc.ExtractLines(data) {
		color := r.FindStringSubmatch(line)[1]
		for _, ic := range ir.FindAllStringSubmatch(line, -1) {
			if containedin[ic[2]] == nil {
				containedin[ic[2]] = make(map[string]bool, 0)
			}
			containedin[ic[2]][color] = true
		}
	}

	contains := make(map[string]bool, 0)
	var check func(color string)
	check = func(color string) {
		for k, _ := range containedin[color] {
			contains[k] = true
			check(k)
		}
	}
	check("shiny gold")
	fmt.Println(len(contains))
}
