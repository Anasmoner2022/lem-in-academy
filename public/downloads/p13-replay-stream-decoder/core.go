package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement replay decoder")

type Room struct {
	Name string
	X, Y int
}
type Snapshot struct {
	Turn      int
	Positions []int
}
type Link struct{ A, B int }
type Replay struct {
	Rooms      []Room
	Links      []Link
	Start, End int
	Frames     []Snapshot
}

func Decode(r io.Reader, maxBytes int64) (Replay, error) { // TODO: framing, grammar, independent validation and immutable snapshots
	return Replay{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: stdin/file boundary and text timeline
	return ErrTODO
}
