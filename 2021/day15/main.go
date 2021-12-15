package main

import (
	"container/heap"
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type pos struct {
	x int
	y int
}

var adjacent = []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	var maxX, maxY = 0, 0

	cave := map[pos]int{}
	for x, l := range lines {
		for y, v := range aoc.DigitList(l) {
			cave[pos{x: x, y: y}] = v
			if x == len(lines)-1 && y == len(aoc.DigitList(l))-1 {
				maxX = x
				maxY = y
			}
		}
	}

	largeCave := map[pos]int{}
	// Create a larger cave.
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for k, v := range cave {
				newRisk := v + x + y
				if newRisk > 9 {
					newRisk = newRisk - 9
				}
				largeCave[pos{
					x: (k.x) + (maxX+1)*x,
					y: (k.y) + (maxY+1)*y,
				}] = newRisk
			}
		}
	}

	// Bump grid to correct size.
	maxX = (maxX+1)*5 - 1
	maxY = (maxY+1)*5 - 1

	riskAt := make(map[pos]int)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(item{pos: pos{0, 0}, riskLevel: 0})

	for pq.Len() > 0 {
		head := heap.Pop(&pq).(item)
		for _, a := range adjacent {
			next := pos{head.pos.x + a.x, head.pos.y + a.y}

			// Check that we are in bounds of the grid.
			if next.x > maxX || next.x < 0 || next.y > maxY || next.y < 0 {
				continue
			}

			nextRisk := head.riskLevel + largeCave[next]
			if sAt, ok := riskAt[next]; ok && sAt <= nextRisk {
				continue
			}

			riskAt[next] = nextRisk
			pq.Push(item{pos: next, riskLevel: nextRisk})
		}
	}

	fmt.Println(riskAt[pos{(maxX), (maxY)}])
}

type item struct {
	pos       pos
	riskLevel int
}

type PriorityQueue []item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].riskLevel < pq[j].riskLevel
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
