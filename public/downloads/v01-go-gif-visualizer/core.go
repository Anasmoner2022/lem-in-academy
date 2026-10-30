package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement visualizer integration")

type Options struct {
	Path          string
	Width, Height int
	Overwrite     bool
}

func Visualize(input io.Reader, options Options) error { // TODO: validated replay, coordinate transform, bounded GIF and companion files
	return ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: flags, stdin, safe output finalization and summary
	return ErrTODO
}
