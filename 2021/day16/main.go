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

	_, value := parse(hexToBinary(lines[0]), 0)
	fmt.Println("res", value)
}

func parse(data string, offset int) (int, int) {
	fmt.Printf("parsing data %d from offset %d\n", len(data), offset)
	fmt.Println(data[offset:])
	v := binaryToDecimal(data[offset+0 : offset+3])
	pkg := binaryToDecimal(data[offset+3 : offset+3+3])

	fmt.Println("version", v, "package", pkg)
	// literal value cant have any nested packets just return version and new offset.
	if pkg == 4 {
		offset += 6
		for {
			last := data[offset] == '0'
			offset += 5
			if last {
				break
			}
		}
		return offset, v
	}
	lt := data[offset+6]
	if lt == '0' {
		l := binaryToDecimal(data[offset+7 : offset+7+15])
		offset = offset + 7 + 15
		end := offset + int(l)
		for {
			nextOffset, addV := parse(data, offset)
			v += addV
			offset = nextOffset
			fmt.Println(end, nextOffset, end)
			if nextOffset >= end {
				break
			}
		}
	} else {
		l := binaryToDecimal(data[offset+7 : offset+7+11])
		offset = offset + 7 + 11
		for i := 0; i < l; i++ {
			fmt.Println(offset)
			nextOffset, addV := parse(data, offset)
			v += addV
			offset = nextOffset
		}
	}
	return offset, v
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
