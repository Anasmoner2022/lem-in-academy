package main

import (
	"strings"
	"testing"
)

func TestParseAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing parser; add the complete mutation table")
	c, e := Parse(strings.Split("3\n##start\ns 0 0\nm 1 0\n##end\nt 2 0\ns-m\nm-t", "\n"))
	if e != nil || c.Ants != 3 || len(c.Rooms) != 3 || len(c.Links) != 2 {
		t.Fatalf("%+v %v", c, e)
	}
}
