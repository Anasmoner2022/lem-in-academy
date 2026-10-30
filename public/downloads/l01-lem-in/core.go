package main

import (
	"errors"
	"io"
)

var ErrTODO = errors.New("TODO: implement lem-in integration")

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
	Echo       string
}
type Move struct{ Ant, To int }
type Schedule struct{ Turns [][]Move }

func ReadColony(r io.Reader) (Colony, error) { // TODO: bounded reading, grammar and complete validation
	return Colony{}, ErrTODO
}
func Optimize(c Colony) (Schedule, error) { // TODO: exact deadline planning using prepared components
	return Schedule{}, ErrTODO
}
func ValidateSchedule(c Colony, s Schedule) error { // TODO: independent movement replay and completion
	return ErrTODO
}
func FormatTranscript(c Colony, s Schedule) ([]byte, error) { // TODO: normalized echo, separator, deterministic move tokens
	return nil, ErrTODO
}
func Execute(args []string, out, errout io.Writer) int { // TODO: process adapter with typed errors, streams and exit statuses
	_, _ = io.WriteString(errout, ErrTODO.Error()+"\n")
	return 3
}
