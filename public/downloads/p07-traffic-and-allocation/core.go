package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement allocation")

type Allocation struct {
	Turns  int
	Counts []int
}

func Allocate(lengths []int, ants int) (Allocation, error) { // TODO: monotone deadline capacity and deterministic counts
	return Allocation{}, ErrTODO
}

type Path struct{ Rooms []int }
type Move struct{ Ant, To int }

func Pipeline(paths []Path, counts []int) ([][]Move, error) { // TODO: validate internal disjointness, assign IDs and generate turn batches
	return nil, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: decode lengths, calculate and print
	return ErrTODO
}
