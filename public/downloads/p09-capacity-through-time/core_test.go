package main

import "testing"

func TestExpansionAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing model; add flow feasibility tests using P08")
	m, e := Expand(Graph{Adj: [][]int{{1}, {0}}}, 0, 1, 3, 1, Budget{1000, 10000})
	if e != nil || m.Vertices != 10 || len(m.Arcs) == 0 {
		t.Fatalf("%+v %v", m, e)
	}
}
