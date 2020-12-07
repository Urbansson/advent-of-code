package main

import (
	"fmt"
	"regexp"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	containedin := make(map[string][]struct {
		ct  int
		col string
	}, 0)
	r, _ := regexp.Compile("(.+?) bags contain")
	ir, _ := regexp.Compile("(\\d+) (.+?) bags?[,.]")
	for _, line := range aoc.ExtractLines(data) {
		color := r.FindStringSubmatch(line)[1]
		for _, ic := range ir.FindAllStringSubmatch(line, -1) {
			containedin[color] = append(containedin[color], struct {
				ct  int
				col string
			}{ct: aoc.Atoi(ic[1]), col: ic[2]})
		}
	}
	var check func(color string) int
	check = func(color string) int {
		count := 0
		for _, v := range containedin[color] {
			count += v.ct
			count += v.ct * check(v.col)
		}
		return count
	}
	fmt.Println(check("shiny gold"))
}
