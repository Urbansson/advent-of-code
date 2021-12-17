package main

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	fmt.Println(Parse(hexToBinary(lines[0])))
}

func Parse(data string) (int, int) {
	_, versions, value := parse(data, 0)
	return versions, value
}

// parse returns offset, version sum, value
func parse(data string, offset int) (int, int, int) {
	v := binaryToDecimal(data[offset+0 : offset+3])
	pkg := binaryToDecimal(data[offset+3 : offset+3+3])
	// literal value cant have any nested packets just return new offset, version and value.
	if pkg == 4 {
		offset += 6
		value := ""
		for {
			last := data[offset] == '0'
			value += data[offset+1 : offset+5]
			offset += 5
			if last {
				break
			}
		}
		return offset, v, binaryToDecimal(value)
	}
	values := []int{}
	lt := data[offset+6]
	if lt == '0' {
		l := binaryToDecimal(data[offset+7 : offset+7+15])
		offset = offset + 7 + 15
		end := offset + int(l)
		for {
			nextOffset, addV, value := parse(data, offset)
			v += addV
			offset = nextOffset
			values = append(values, value)
			if nextOffset >= end {
				break
			}
		}
	} else {
		l := binaryToDecimal(data[offset+7 : offset+7+11])
		offset = offset + 7 + 11
		for i := 0; i < l; i++ {
			nextOffset, addV, value := parse(data, offset)
			v += addV
			offset = nextOffset
			values = append(values, value)
		}
	}

	res := 0
	switch pkg {
	case 0:
		for _, v := range values {
			res += v
		}
	case 1:
		for _, v := range values {
			res *= v
		}
	case 2:
		for _, v := range values {
			res = aoc.Min(res, v)
		}
	case 3:
		for _, v := range values {
			res = aoc.Max(res, v)
		}
	case 5:
		if values[0] > values[1] {
			res = 1
		}
	case 6:
		if values[0] < values[1] {
			res = 1
		}
	case 7:
		if values[0] == values[1] {
			res = 1
		}
	}
	return offset, v, res
}

func hexToBinary(s string) string {
	d, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	res := ""
	for _, d := range d {
		res += fmt.Sprintf("%08b", d)
	}
	return res
}

func binaryToDecimal(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
