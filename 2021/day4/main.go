package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type pos struct {
	x int
	y int
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	numbers := aoc.ExtractInts(lines[0])
	boards := []map[pos]int{}

	currBoard := map[pos]int{}
	for x, l := range lines[2:] {
		if len(l) == 0 {
			fmt.Println("empty")
			boards = append(boards, currBoard)
			currBoard = make(map[pos]int)
			continue
		}
		currBoardRow := x % 6
		for y, num := range aoc.ExtractInts(l) {
			currBoard[pos{currBoardRow, y}] = num
		}
	}
	boards = append(boards, currBoard)

	marked := make(map[int]bool)
	for _, n := range numbers {
		marked[n] = true
		for _, b := range boards {
			if bingo(b, marked) {
				fmt.Println(sum(b, marked) * n)
				return
			}
		}
	}
	fmt.Println("hmm", len(boards))
}

func sum(board map[pos]int, marked map[int]bool) int {
	sum := 0
	for _, v := range board {
		if _, ok := marked[v]; !ok {
			sum += v
		}
	}
	return sum
}

func bingo(board map[pos]int, marked map[int]bool) bool {
	x, y := [5]int{}, [5]int{}
	for p, v := range board {
		if _, ok := marked[v]; ok {
			x[p.x]++
			y[p.y]++
			if x[p.x] == 5 || y[p.y] == 5 {
				return true
			}
		}
	}
	return false
}
