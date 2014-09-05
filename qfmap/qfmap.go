package qfmap

type UnionFind struct {
	id    map[int]int
	next	int
}

func New() *UnionFind {
	return &UnionFind{id: make(map[int]int)}
}

func (uf UnionFind) Find(p int) int {
	return uf.id[p]
}

func (uf UnionFind) Connected(p, q int) bool {
  p_id, p_in_map := uf.id[p]
  q_id, q_in_map := uf.id[q]

	if !(p_in_map && q_in_map) { return false }

	return p_id == q_id
}

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
	
	if next_consumed { uf.next++ }
	
  for i, n := range uf.id {
		if n == p_id {
			uf.id[i] = q_id
		}
	}
}
