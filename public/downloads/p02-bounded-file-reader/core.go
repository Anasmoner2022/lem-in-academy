package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement bounded reader")

type Limits struct{ MaxBytes, MaxLine int }
type Document struct {
	Lines []string
	Echo  string
}

func ReadDocument(r io.Reader, limits Limits) (Document, error) { // TODO: bounds, line normalization and reader errors
	return Document{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: file ownership, report and writer errors
	return ErrTODO
}
