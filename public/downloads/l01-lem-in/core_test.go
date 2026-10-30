package main

import (
	"strings"
	"testing"
)

func TestCorridorAcceptance(t *testing.T) {
	t.Skip("TODO: enable after implementation; add full CLI, oracle and movement tests")
	c, e := ReadColony(strings.NewReader("3\n##start\ns 0 0\na 1 0\n##end\nt 2 0\ns-a\na-t\n"))
	if e != nil {
		t.Fatal(e)
	}
	s, e := Optimize(c)
	if e != nil || len(s.Turns) != 4 {
		t.Fatalf("%+v %v", s, e)
	}
	if e = ValidateSchedule(c, s); e != nil {
		t.Fatal(e)
	}
}
