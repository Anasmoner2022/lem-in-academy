package main

import "testing"

func TestOracleAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing oracle; add graph enumeration and transcript mutations")
	turns, e := Oracle(Graph{Adj: [][]int{{1}, {0, 2}, {1}}}, 0, 2, 3, 10000)
	if e != nil || turns != 4 {
		t.Fatalf("%d %v", turns, e)
	}
}
func TestDifferential(t *testing.T) {
	t.Skip("TODO: implement the 3294-case campaign and print counts")
}
func FuzzTranscript(f *testing.F) {
	f.Add("not a transcript")
	f.Fuzz(func(t *testing.T, s string) { t.Skip("TODO: bound input, call checker and assert no panic") })
}
