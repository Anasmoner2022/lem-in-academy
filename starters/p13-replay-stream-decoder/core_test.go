package main

import (
	"strings"
	"testing"
)

func TestReplayAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing decoder; add pipe and aliasing cases")
	s := "2\n##start\ns 0 0\na 1 0\n##end\nt 2 0\ns-a\na-t\n\nL1-a\nL1-t L2-a\nL2-t\n"
	r, e := Decode(strings.NewReader(s), 1<<20)
	if e != nil || len(r.Frames) != 4 {
		t.Fatalf("%+v %v", r, e)
	}
}
