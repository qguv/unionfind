package qfmap

import "testing"

func TestNew(t *testing.T) {
	New()
}

func TestUnion(t *testing.T) {
	var uf *UnionFind

	uf = New()
	uf.Union(3, 4)
	uf.Union(1, 3)

	uf = New()
	uf.Union(0, 1)
	uf.Union(2, 1)
}

func TestConnected(t *testing.T) {
	uf := New()
	uf.Union(3, 4)
	uf.Union(1, 3)
	if !uf.Connected(1, 3) {
		t.Error("connected gives false negative on input union")
	}
	if !uf.Connected(1, 4) {
		t.Error("connected gives false negative on transitive union")
	}
	if uf.Connected(6, 7) {
		t.Error("connected gives false positive on random nodes")
	}
	if uf.Connected(3, 7) {
		t.Error("connected gives false positive on reassigned nodes")
	}
	if uf.Connected(1, 7) {
		t.Error("connected gives false positive on reassigned nodes")
	}
}
