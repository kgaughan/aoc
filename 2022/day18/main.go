package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type voxel struct {
	x, y, z int
}

type adjacency struct {
	top, bottom, left, right, front, back bool
}

func (a adjacency) Exposed() int {
	exposed := 0
	if !a.left {
		exposed++
	}
	if !a.right {
		exposed++
	}
	if !a.top {
		exposed++
	}
	if !a.bottom {
		exposed++
	}
	if !a.back {
		exposed++
	}
	if !a.front {
		exposed++
	}
	return exposed
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	surface := make(map[voxel]*adjacency)
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	minZ, maxZ := math.MaxInt, math.MinInt
	for {
		var c voxel
		if _, err := fmt.Fscanf(f, "%d,%d,%d\n", &c.x, &c.y, &c.z); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		surface[c] = &adjacency{}
		minX = min(c.x, minX)
		minY = min(c.y, minY)
		minZ = min(c.z, minZ)
		maxX = max(c.x, maxX)
		maxY = max(c.y, maxY)
		maxZ = max(c.z, maxZ)
	}

	// Now we run over all the voxels, figuring out if there's one adjacent to
	// the right, bottom, or back.
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				self, exists := surface[voxel{x: x, y: y, z: z}]
				if !exists {
					continue
				}
				if right, exists := surface[voxel{x: x + 1, y: y, z: z}]; exists {
					right.left = true
					self.right = true
				}
				if bottom, exists := surface[voxel{x: x, y: y + 1, z: z}]; exists {
					bottom.top = true
					self.bottom = true
				}
				if back, exists := surface[voxel{x: x, y: y, z: z + 1}]; exists {
					back.front = true
					self.back = true
				}
			}
		}
	}
	area := 0
	for _, adjacency := range surface {
		area += adjacency.Exposed()
	}
	fmt.Printf("Part 1: %v exposed voxel surfaces\n", area)
}
