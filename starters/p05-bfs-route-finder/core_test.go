package main

import "testing"

func TestShortestAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing BFS; add cycle, queue and unreachable cases")
	g := Graph{Names: []string{"s", "a", "t"}, Adj: [][]int{{1}, {0, 2}, {1}}}
	r, e := Shortest(g, 0, 2)
	if e != nil || r.Distance != 2 || len(r.Vertices) != 3 {
		t.Fatalf("%+v %v", r, e)
	}
}
