package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	size           int
	subdirectories map[string]*Directory
}

func NewDirectory() *Directory {
	return &Directory{
		size:           0,
		subdirectories: make(map[string]*Directory),
	}
}

func (d Directory) TotalSize() int {
	result := d.size
	// This is wasteful. We could cache the result, but that's messy.
	for _, dir := range d.subdirectories {
		result += dir.TotalSize()
	}
	return result
}

func (d *Directory) AddFile(path []string, file string, size int) {
	if len(path) == 0 {
		d.size += size
	} else {
		subdir, exists := d.subdirectories[path[0]]
		if !exists {
			subdir = NewDirectory()
			d.subdirectories[path[0]] = subdir
		}
		subdir.AddFile(path[1:], file, size)
	}
}

func (d *Directory) String() string {
	buf := &strings.Builder{}
	d.describe(buf, "/", "")
	return buf.String()
}

func (d *Directory) describe(buf *strings.Builder, name, indent string) {
	buf.WriteString(fmt.Sprintf("%s- %s (dir, files=%v)\n", indent, name, d.TotalSize()))
	for name, dir := range d.subdirectories {
		dir.describe(buf, name, indent+"  ")
	}
}

func (d *Directory) Walk(fn func(string, int)) {
	for name, dir := range d.subdirectories {
		fn(name, dir.TotalSize())
		dir.Walk(fn)
	}
}

const (
	P1_MAX       = 100000
	P2_TARGET    = 30000000
	P2_DISCSPACE = 70000000
)

var data = flag.String("data", "input.txt", "Name of input file")

func main() {
	flag.Parse()
	f, err := os.Open(*data)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cwd := make([]string, 0, 10)
	root := NewDirectory()

	// Parse
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		switch parts[0] {
		case "$":
			switch parts[1] {
			case "cd":
				switch parts[2] {
				case "/":
					cwd = cwd[:0]
				case "..":
					cwd = cwd[:len(cwd)-1]
				default:
					cwd = append(cwd, parts[2])
				}
			case "ls":
				// ignore
			}
		case "dir":
			// ignore
		default:
			n, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("Could not parts %q as an int: %v", parts[0], err)
			}
			root.AddFile(cwd, parts[1], n)
		}
	}

	// The amount of free space we need to free up
	p2Target := P2_TARGET - (P2_DISCSPACE - root.TotalSize())

	p1Answer := 0 // Part 1: find directories with a total size of at most P1_MAX
	p2Answer := 0 // Part 2: find the smallest directory that's at least the larget
	root.Walk(func(name string, size int) {
		if size <= P1_MAX {
			p1Answer += size
		}
		// For part 2, we want the smallest directory larger than the target,
		// but will settle for the largest directory less than the target.
		if size >= p2Target && (size < p2Answer || p2Answer == 0) {
			p2Answer = size
		}
	})
	fmt.Printf("Part 1: %v\n", p1Answer)
	fmt.Printf("Part 2: %v\n", p2Answer)
}
