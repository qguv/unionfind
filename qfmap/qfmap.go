package qfmap

/*
UnionFind contains information about how given numbers are related. Use
UnionFind.Union(2, 3) to specify that 2 and 3 belong to the same set.
Additional calls to Union can create relationships by the transitive property.
Query the relationship between numbers with the Connected or Find methods.

This implementation of UnionFind uses quick-find and stores connection data in
a map.
*/
type UnionFind struct {
	id   map[int]int
	next int
}

// New initializes a new UnionFind struct and returns a pointer.
func New() *UnionFind {
	return &UnionFind{id: make(map[int]int)}
}

// Find returns the set ID of a given element.
func (uf UnionFind) Find(p int) int {
	return uf.id[p]
}

// Connected determines whether two elements belong to the same set.
func (uf UnionFind) Connected(p, q int) bool {
	p_id, p_in_map := uf.id[p]
	q_id, q_in_map := uf.id[q]
	if !(p_in_map && q_in_map) {
		return false
	}
	return p_id == q_id
}

// Union specifies two elements as belonging to the same set for future query
// with the Connected or Find methods.
func (uf *UnionFind) Union(p, q int) {
	if uf.Connected(p, q) {
		return
	}

	p_id, p_in_map := uf.id[p]
	q_id, q_in_map := uf.id[q]

	var next_consumed bool

	if !p_in_map {
		uf.id[p] = uf.next
		p_id = uf.next
		next_consumed = true
	}

	if !q_in_map {
		uf.id[q] = uf.next
		q_id = uf.next
		next_consumed = true
	}

	if next_consumed {
		uf.next++
	}

	for i, n := range uf.id {
		if n == p_id {
			uf.id[i] = q_id
		}
	}
}
