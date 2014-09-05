package qfarray

type UnionFind struct {
	id    []int
	count int
}

func New(count int) *UnionFind {
	id := make([]int, count)
	for i := 0; i < count; i++ {
		id[i] = i
	}
	return &UnionFind{id: id, count: count}
}

func (uf UnionFind) Find(p int) int {
	return uf.id[p]
}

func (uf UnionFind) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

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
