package day9

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

func (es Edges) Distance(chain []string) (int, bool) {
	total := 0
	for i := 0; i < len(chain)-1; i++ {
		if dist, ok := es.Get(chain[i], chain[i+1]); ok {
			total += dist
		} else {
			return 0, false
		}
	}
	return total, true
}
