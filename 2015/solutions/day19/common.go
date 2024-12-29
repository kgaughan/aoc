package day19

import (
	"strings"
)

func allIndexes(str, substr string) []int {
	indexes := make([]int, 0, 10)
	remaining := str
	offset := 0
	for {
		i := strings.Index(remaining, substr)
		if i == -1 {
			break
		}
		indexes = append(indexes, offset+i)
		offset += i + len(substr)
		remaining = remaining[i+len(substr):]
	}
	return indexes
}

type Item struct {
	steps    int
	molecule string
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if len(pq[i].molecule) > len(pq[j].molecule) {
		return true
	}
	return pq[i].steps > pq[j].steps
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
