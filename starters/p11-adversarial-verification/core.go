package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement independent verification")
var ErrBudget = errors.New("oracle budget exhausted")

type Graph struct{ Adj [][]int }
type Report struct{ Turns, Arrived int }

func CheckTranscript(original []byte, transcript io.Reader) (Report, error) { // TODO: bounded echo/token parser and independent replay
	return Report{}, ErrTODO
}
func Oracle(g Graph, start, end, ants, maxStates int) (int, error) { // TODO: exhaustive occupancy-state BFS, no flow engine calls
	return 0, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: run tiny fixture and report optimum
	return ErrTODO
}
