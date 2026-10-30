package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement colony parser")

type Room struct {
	Name string
	X, Y int
}
type Link struct{ A, B int }
type Colony struct {
	Ants       int
	Rooms      []Room
	Links      []Link
	Start, End int
}

func Parse(lines []string) (Colony, error) { // TODO: phases, commands, rooms, terminals, unique links
	return Colony{}, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: reuse bounded reader, parse, inspect
	return ErrTODO
}
