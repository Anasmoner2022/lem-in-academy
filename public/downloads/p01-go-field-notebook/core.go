package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement notebook")

type Room struct {
	Name string
	X, Y int
}

func ValidateRooms(rooms []Room) error { // TODO: name and uniqueness invariants
	return ErrTODO
}
func Inventory(rooms []Room) (string, error) { // TODO: validate, stable ordering, no caller mutation
	return "", ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: exactly one file, decode, validate and write
	return ErrTODO
}
