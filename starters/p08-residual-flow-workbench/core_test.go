package main

import "testing"

func TestFlowAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing flow; add reroute and conservation cases")
	r, e := Solve(4, []EdgeSpec{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 3, 1}}, 0, 3, 2)
	if e != nil || r.Value != 2 {
		t.Fatalf("%+v %v", r, e)
	}
}
