package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement graph inventory")

type Pair struct{ A, B string }
type Graph struct {
	Names []string
	Adj   [][]int
}

func Build(names []string, links []Pair) (Graph, error) { // TODO: canonical IDs and undirected adjacency
	return Graph{}, ErrTODO
}
func Components(g Graph) ([][]int, error) { // TODO: iterative traversal including isolated vertices
	return nil, ErrTODO
}
func HasCycle(g Graph) (bool, error) { // TODO: parent-aware undirected traversal
	return false, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: decode fixture, inspect graph, stable report
	return ErrTODO
}
