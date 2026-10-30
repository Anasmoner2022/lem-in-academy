package main

import "testing"

func TestInventoryAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing Inventory; add validation and I/O cases")
	got, err := Inventory([]Room{{Name: "z", X: 1}, {Name: "a", X: -2, Y: 5}})
	if err != nil || got != "a -2 5\nz 1 0\n" {
		t.Fatalf("got %q, %v", got, err)
	}
}
