package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type Node struct {
	parent *Node
	size   int
	isDir  bool
	name   string
	files  []*Node
}

func (f *Node) addFile(name string, size int) {
	f.files = append(f.files, &Node{name: name, parent: f})
	f.increaseSize(size)
}

func (f *Node) addDir(name string) {
	f.files = append(f.files, &Node{name: name, parent: f, files: []*Node{}, isDir: true})
}

func (f *Node) increaseSize(size int) {
	f.size += size
	if f.parent != nil {
		f.parent.increaseSize(size)
	}
}

func (f *Node) findDir(name string) *Node {
	for _, d := range f.files {
		if d.name == name && d.isDir {
			return d
		}
	}
	return nil
}

func main() {
	data := aoc.ReadStdinRaw()
	lines := aoc.ExtractLines(data)

	root := &Node{name: "/", files: []*Node{}}
	current := root

	for i, line := range lines {
		switch {
		case line == "$ cd /":
			current = root
		case line == "$ cd ..":
			current = current.parent
		case strings.HasPrefix(line, "$ cd "):
			pieces := strings.Split(line, " ")
			current = current.findDir(pieces[2])
		case line == "$ ls":
			for {
				// Peek, stop if we find another command or reached end of input
				if i == len(lines)-1 || lines[i+1][0] == '$' {
					break
				}
				// Consume next line
				i++
				line = lines[i]
				pieces := strings.Split(line, " ")
				// Add file or dir
				if pieces[0] == "dir" {
					current.addDir(pieces[1])
				} else {
					size, _ := strconv.Atoi(pieces[0])
					current.addFile(pieces[1], size)
				}
			}
		}
	}

	fmt.Println(findCandidatesFileSize(root, 100000))
}

func findCandidatesFileSize(dir *Node, minFileSize int) int {
	total := 0

	if dir.size <= minFileSize {
		total += dir.size
	}

	for _, file := range dir.files {
		if file.isDir {
			total += findCandidatesFileSize(file, minFileSize)
		}
	}

	return total
}
