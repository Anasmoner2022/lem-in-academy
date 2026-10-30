package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement BFS")
var ErrUnreachable = errors.New("unreachable")

type Graph struct {
	Names []string
	Adj   [][]int
}
type Route struct {
	Vertices []int
	Distance int
}

func Shortest(g Graph, start, end int) (Route, error) { // TODO: queue, discovery, parents and guarded reconstruction
	return Route{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: reuse graph construction and print route
	return ErrTODO
}
