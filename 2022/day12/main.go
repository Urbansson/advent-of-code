package main

import (
	"fmt"
	"math"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	start := aoc.XY{X: 0, Y: 0}
	end := aoc.XY{X: 0, Y: 0}

	area := map[aoc.XY]int{}
	for x, l := range lines {
		for y, v := range l {
			area[aoc.XY{X: x, Y: y}] = int(v)
			if v == 'S' {
				start = aoc.XY{X: x, Y: y}
				area[aoc.XY{X: x, Y: y}] = int('a')
			}
			if v == 'E' {
				end = aoc.XY{X: x, Y: y}
				area[aoc.XY{X: x, Y: y}] = int('z')
			}
		}
	}

	fmt.Println(start, end)

	// Print the hightmap
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Print(string(area[aoc.XY{X: x, Y: y}]))
		}
		fmt.Println()
	}

	//Find the shortest path from start to end using bfs.
	res := bfs(area, start, end)
	fmt.Println(res)
}

func inBound(pos aoc.XY, grid map[aoc.XY]int) bool {
	maxX := 8
	maxY := 5

	return 0 <= pos.X && pos.X < maxX && 0 <= pos.Y && pos.Y < maxY
}

var adjacent = []aoc.XY{{X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: -1, Y: 0}}

func bfs(grid map[aoc.XY]int, start aoc.XY, end aoc.XY) int {
	var q PriorityQueue[item]
	q.PushBack(item{start, 0})
	seen := make(map[aoc.XY]interface{})
	for q.Len() > 0 {
		v := q.Pop()
		pos, dist := v.XY, v.dist

		if pos.X == end.X && pos.Y == end.Y {
			return dist
		}
		if _, ok := seen[pos]; ok {
			continue
		}
		seen[pos] = true
		x, y := pos.X, pos.Y
		for _, coord := range adjacent {
			dx, dy := coord.X, coord.Y
			if inBound(aoc.XY{X: x + dx, Y: y + dy}, grid) && grid[aoc.XY{X: x + dx, Y: y + dy}]-grid[aoc.XY{X: x, Y: y}] <= 1 {
				q.PushBack(item{
					XY: aoc.XY{
						X: x + dx, Y: y + dy,
					},
					dist: dist + 1,
				})
			}
		}

	}
	return math.MaxInt
}

type item struct {
	aoc.XY
	dist int
}

type PriorityQueue[C any] []C

func (pq *PriorityQueue[C]) Len() int {
	return len(*pq)
}

func (pq PriorityQueue[C]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[C]) Push(x C) {
	*pq = append(*pq, x)
}

func (pq *PriorityQueue[C]) PushBack(x C) {
	*pq = append([]C{x}, *pq...)
}

func (pq *PriorityQueue[C]) Pop() C {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[C]) PopBack() C {
	old := *pq
	item := old[0]
	*pq = old[1:]
	return item
}

/*
func (pq *PriorityQueue) Len() int {
	return len(pq)
}

/*
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x C) {
	item := x.(item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) PushBack(x C) {
	i := x.(item)
	*pq = append([]item{i}, *pq...)
}

func (pq *PriorityQueue) Pop() C {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) PopBack() C {
	old := *pq
	item := old[0]
	*pq = old[1:]
	return item
}
*/
