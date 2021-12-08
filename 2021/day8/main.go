package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	res := 0
	for _, line := range lines {
		p := strings.Split(line, " | ")
		inputs := strings.Split(p[0], " ")
		outputs := strings.Split(p[1], " ")
		wires := [10]string{}
		// Find known inputs.
		for _, v := range inputs {
			l := len(v)
			if l == 2 {
				wires[1] = v
			} else if l == 3 {
				wires[7] = v
			} else if l == 4 {
				wires[4] = v
			} else if l == 7 {
				wires[8] = v
			}
		}
		//Find the remaining ones
		for i := 0; i < 10; i++ {
			for _, v := range inputs {
				if wires[2] == "" && len(v) == 5 && overlaps(v, wires[7]) == 2 && overlaps(v, wires[1]) == 1 && overlaps(v, wires[3]) == 4 && overlaps(v, wires[4]) == 2 {
					wires[2] = v
				}
				if wires[3] == "" && len(v) == 5 && overlaps(v, wires[1]) == 2 {
					wires[3] = v
				}
				if wires[5] == "" && len(v) == 5 && overlaps(v, wires[2]) == 3 && overlaps(v, wires[4]) == 3 {
					wires[5] = v
				}
				if wires[6] == "" && len(v) == 6 && overlaps(v, wires[1]) == 1 && overlaps(v, wires[2]) == 4 {
					wires[6] = v
				}
				if wires[9] == "" && len(v) == 6 && overlaps(v, wires[7]) == 3 && overlaps(v, wires[0]) == 5 {
					wires[9] = v
				}
				if wires[0] == "" && len(v) == 6 && overlaps(v, wires[1]) == 2 && overlaps(v, wires[2]) == 4 && overlaps(v, wires[3]) == 4 {
					wires[0] = v
				}
			}
		}
		var o string
		for k, v := range wires {
			wires[k] = aoc.SortString(v)
		}
		for _, v := range outputs {
			v = aoc.SortString(v)
			for k, w := range wires {
				if w == v {
					o += fmt.Sprintf("%d", k)
				}
			}
		}
		res += aoc.Atoi(o)
	}
	fmt.Println(res)
}

// overlaps, return number number of overlapping segments in the display.
func overlaps(a, b string) int {
	count := 0
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			count++
		}
	}
	return count
}
