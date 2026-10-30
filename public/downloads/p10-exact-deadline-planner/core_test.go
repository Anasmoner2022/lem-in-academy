package main

import "testing"

func TestPlanAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing planner; validate movements independently")
	p, e := Plan(Graph{Adj: [][]int{{1}, {0, 2}, {1}}}, 0, 2, 3)
	if e != nil || len(p.Turns) != 4 {
		t.Fatalf("%+v %v", p, e)
	}
}
