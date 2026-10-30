package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement release boundaries")

type Counts struct{ Vertices, OriginalArcs, ResidualArcs int64 }

func Estimate(vertices, edges, horizon int64) (Counts, error) { // TODO: checked arithmetic from published formulas
	return Counts{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: decode count fixture and format checked estimates
	return ErrTODO
}
