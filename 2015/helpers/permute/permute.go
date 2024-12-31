package permute

type Permute struct {
	source []string
	idx    []int
	i      int
}

func NewPermute(entries []string) *Permute {
	source := make([]string, len(entries))
	copy(source, entries)

	return &Permute{
		source: source,
		idx:    make([]int, len(entries)),
		i:      1,
	}
}

func (p *Permute) Get() ([]string, bool) {
	// Heap's algorithm for generating permutations:
	// https://en.wikipedia.org/wiki/Heap%27s_algorithm
	for p.i < len(p.source) {
		if p.idx[p.i] < p.i {
			swap := p.i % 2 * p.idx[p.i]
			p.source[swap], p.source[p.i] = p.source[p.i], p.source[swap]
			p.idx[p.i]++
			p.i = 1
			return p.source, false
		}
		p.idx[p.i] = 0
		p.i++
	}
	return nil, true
}
