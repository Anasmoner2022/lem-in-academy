package main

import "testing"

func TestEstimateAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing budgets; add overflow and refusal cases")
	c, e := Estimate(3, 2, 3)
	if e != nil || c.Vertices != 36 || c.OriginalArcs != 51 || c.ResidualArcs != 102 {
		t.Fatalf("%+v %v", c, e)
	}
}
