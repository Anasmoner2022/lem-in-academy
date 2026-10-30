package main

import "testing"

func TestPipelineAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing Step; add all rejection rules")
	g := Graph{Adj: [][]int{{1}, {0, 2}, {1}}}
	next, e := Step(g, 0, 2, []int{1, 0}, []Move{{1, 2}, {2, 1}})
	if e != nil || len(next) != 2 || next[0] != 2 || next[1] != 1 {
		t.Fatalf("%v %v", next, e)
	}
}
