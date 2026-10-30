package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement exact planner")

type Graph struct{ Adj [][]int }
type Move struct{ Ant, To int }
type PlanResult struct {
	Turns        [][]Move
	LowerBound   int
	PreviousFlow int
}

func Plan(g Graph, start, end, ants int) (PlanResult, error) { // TODO: bounds, exact deadline probes, integral extraction and independent validation
	return PlanResult{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: decode validated graph fixture and format plan
	return ErrTODO
}
