package graph

type edge struct {
	from, to string
}

type Edges map[edge]int

func (es Edges) Add(first, second string, dist int) {
	es[edge{first, second}] = dist
}

func (es Edges) Get(first, second string) (int, bool) {
	if edge, ok := es[edge{first, second}]; ok {
		return edge, true
	}
	if edge, ok := es[edge{second, first}]; ok {
		return edge, true
	}
	return 0, false
}

func (es Edges) Distance(start string, rest []string) int {
	total, _ := es.Get(start, rest[0])
	for i := 0; i < len(rest)-1; i++ {
		dist, _ := es.Get(rest[i], rest[i+1])
		total += dist
	}
	return total
}
