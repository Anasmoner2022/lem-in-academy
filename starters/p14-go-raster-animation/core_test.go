package main

import "testing"

func TestAnimationAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing renderer; add pixel, transform and decode-back tests")
	a, e := Animate(Scene{}, [][]Pixel{{{20, 20}}, {{80, 20}}}, 100, 60)
	if e != nil || len(a.Image) != 6 || len(a.Delay) != 6 {
		t.Fatalf("%+v %v", a, e)
	}
}
