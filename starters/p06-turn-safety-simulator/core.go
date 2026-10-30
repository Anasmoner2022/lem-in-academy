package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement turn safety")

type Graph struct{ Adj [][]int }
type Move struct{ Ant, To int }

func Step(g Graph, start, end int, positions []int, moves []Move) ([]int, error) { // TODO: identities, edges, departures, arrivals and atomic validation
	return nil, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: decode replay, start positions, validate completion
	return ErrTODO
}
