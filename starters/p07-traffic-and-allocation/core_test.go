package main

import "testing"

func TestAllocationAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing allocation; add capacity and overlap tests")
	a, e := Allocate([]int{2, 4}, 4)
	if e != nil || a.Turns != 4 || len(a.Counts) != 2 || a.Counts[0] != 3 || a.Counts[1] != 1 {
		t.Fatalf("%+v %v", a, e)
	}
}
