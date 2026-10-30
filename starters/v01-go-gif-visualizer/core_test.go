package main

import (
	"image/gif"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestVisualizerAcceptance(t *testing.T) {
	t.Skip("TODO: remove after implementation; add replay, cleanup and pixel checks")
	path := filepath.Join(t.TempDir(), "farm.gif")
	transcript := "2\n##start\ns 0 0\na 1 0\n##end\nt 2 0\ns-a\na-t\n\nL1-a\nL1-t L2-a\nL2-t\n"
	if e := Visualize(strings.NewReader(transcript), Options{Path: path, Width: 320, Height: 240}); e != nil {
		t.Fatal(e)
	}
	f, e := os.Open(path)
	if e != nil {
		t.Fatal(e)
	}
	defer f.Close()
	a, e := gif.DecodeAll(f)
	if e != nil || len(a.Image) != 14 {
		t.Fatalf("frames=%v err=%v", a, e)
	}
}
