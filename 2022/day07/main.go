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

	// Part 1: find directories with a total size of at most 100000
	p1Answer := 0
	root.Walk(func(name string, size int) {
		if size <= 100000 {
			p1Answer += size
		}
	})
	fmt.Printf("Part 1: %v\n", p1Answer)
}
