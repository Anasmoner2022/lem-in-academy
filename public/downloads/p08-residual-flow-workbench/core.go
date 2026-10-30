package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement residual flow")

type Arc struct{ To, Rev, Residual, Original int }
type EdgeSpec struct{ From, To, Capacity int }
type Result struct {
	Value     int
	Flow      []int
	Reachable []bool
	Exhausted bool
}

func Solve(vertices int, edges []EdgeSpec, source, sink, target int) (Result, error) { // TODO: pair arcs, residual BFS, augmentation and certificates
	return Result{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: decode directed network and print flow
	return ErrTODO
}
