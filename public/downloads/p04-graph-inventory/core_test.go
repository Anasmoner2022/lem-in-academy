package main

import "testing"

func TestComponentsAcceptance(t *testing.T) {
	t.Skip("TODO: implement graph API and remove skip; add cycle and ordering cases")
	g, e := Build([]string{"c", "a", "b"}, []Pair{{"a", "b"}})
	if e != nil {
		t.Fatal(e)
	}
	cs, e := Components(g)
	if e != nil || len(cs) != 2 {
		t.Fatalf("%v %v", cs, e)
	}
}
