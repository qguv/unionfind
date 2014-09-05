package qfarray

/* UnionFind contains information about how given numbers are related. Use
UnionFind.Union(2, 3) to specify that 2 and 3 belong to the same set.
Additional calls to Union can create relationships by the transitive
property. Query the relationship between numbers with the Connected or Find
methods.
This implementation of UnionFind uses quick-find and stores connection data in
a map.
*/
type UnionFind struct {
	id    []int
	count int
}

// New initializes a new UnionFind struct and returns a pointer.
func New(count int) *UnionFind {
	id := make([]int, count)
	for i := 0; i < count; i++ {
		id[i] = i
	}
	return &UnionFind{id: id, count: count}
}

// Find returns the set ID of a given element.
func (uf UnionFind) Find(p int) int {
	return uf.id[p]
}

// Connected determines whether two elements belong to the same set.
func (uf UnionFind) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

// Union specifies two elements as belonging to the same set for future query
// with the Connected or Find methods.
func (uf *UnionFind) Union(p, q int) {
	if uf.Connected(p, q) {
		return
	}
	p_id := uf.id[p]
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == p_id {
			uf.id[i] = uf.id[q]
		}
	}
	uf.count--
}
