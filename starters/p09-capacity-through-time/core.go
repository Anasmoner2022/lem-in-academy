package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement time expansion")
var ErrBudget = errors.New("resource limit exceeded")

type Graph struct{ Adj [][]int }
type ArcSpec struct {
	From, To, Capacity int
	Kind               string
	Time, RoomA, RoomB int
}
type Model struct {
	Vertices, Source, Sink int
	Arcs                   []ArcSpec
}
type Budget struct{ Vertices, ResidualArcs int }

func Expand(g Graph, start, end, ants, horizon int, budget Budget) (Model, error) { // TODO: checked counts, room splits, holds and shared tunnel gadgets
	return Model{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: compile fixture and report model inventory
	return ErrTODO
}
