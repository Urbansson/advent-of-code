package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	s := aoc.ReadInput("./input.txt")
	ids := []int{}
	for _, r := range aoc.ExtractLines(s) {
		row := strings.ReplaceAll(r, "B", "1")
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "R", "1")
		row = strings.ReplaceAll(row, "L", "0")

		j, _ := strconv.ParseInt(row, 2, 16)
		ids = append(ids, int(j))
	}
	sort.Ints(ids)
	for i, v := range ids {
		if ids[i+1]-v == 2 {
			fmt.Println(ids[len(ids)-1], v+1)
			break
		}
	}
}
