package main

import (
	"bufio"
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

func (d *Directory) AddFile(path []string, size int) {
	if len(path) == 0 {
		d.size += size
	} else {
		subdir, exists := d.subdirectories[path[0]]
		if !exists {
			subdir = NewDirectory()
			d.subdirectories[path[0]] = subdir
		}
		subdir.AddFile(path[1:], size)
	}
}

func (d *Directory) Walk(fn func(int)) {
	for _, dir := range d.subdirectories {
		fn(dir.TotalSize())
		dir.Walk(fn)
	}
}

const (
	P1_MAX       = 100_000
	P2_TARGET    = 30_000_000
	P2_DISCSPACE = 70_000_000
)

func main() {
	f, err := os.Open("input.txt")
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
			root.AddFile(cwd, n)
		}
	}

	// The amount of free space we need to free up
	p2Target := P2_TARGET - (P2_DISCSPACE - root.TotalSize())

	p1Answer := 0
	p2Answer := root.TotalSize()
	root.Walk(func(size int) {
		if size <= P1_MAX {
			p1Answer += size
		}
		if size >= p2Target && size < p2Answer {
			p2Answer = size
		}
	})
	fmt.Printf("Part 1: %v\n", p1Answer)
	fmt.Printf("Part 2: %v\n", p2Answer)
}
