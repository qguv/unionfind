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

	t.Log("current state of id map")
	for i, n := range uf.id {
		t.Logf("\t%3v: %3v\n", i, n)
	}
}

// A white-box test for a complicated scenario
// init(10); 6-0 3-4 5-8 7-2 2-1 5-7 0-3 4-2
// Everything but nine should be connected
func TestFind(t *testing.T) {
	uf := New()
	uf.Union(6, 0)
	uf.Union(3, 4)
	uf.Union(5, 8)
	uf.Union(7, 2)
	uf.Union(2, 1)
	uf.Union(5, 7)
	uf.Union(0, 3)
	uf.Union(4, 2)

	commonFind := uf.Find(0)
	for i := 1; i < 9; i++ {
		if f := uf.Find(i); f != commonFind {
			t.Errorf(
				"find is incorrect for %v: found %v, should be %v\n",
				i, f, commonFind,
			)
		}
	}

	if f := uf.Find(9); f == commonFind {
		t.Errorf("find gives incorrect %v for 9, which shouldn't be part of the set.\n", f)
	}

	t.Log("current state of id map")
	for i, n := range uf.id {
		t.Logf("\t%3v: %3v\n", i, n)
	}
}
