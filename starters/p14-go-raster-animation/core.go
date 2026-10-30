package main

import (
	"errors"
	"image"
	"image/gif"
	"io"
)

var ErrTODO = errors.New("TODO: implement Go raster animation")

type Point struct{ X, Y int64 }
type Pixel struct{ X, Y int }

func Fit(points []Point, width, height, padding int) ([]Pixel, error) { // TODO: checked aspect-preserving coordinate transform
	return nil, ErrTODO
}

type Scene struct {
	Rooms []Pixel
	Links [][2]int
	Ants  []Pixel
}

func Render(scene Scene, width, height int) (*image.Paletted, error) { // TODO: palette, background, line/disc primitives and bounds
	return nil, ErrTODO
}
func Animate(scene Scene, boundaries [][]Pixel, width, height int) (*gif.GIF, error) { // TODO: budgets, independent frames, interpolation and delays
	return nil, ErrTODO
}
func Run(args []string, input io.Reader, output io.Writer) error { // TODO: JSON fixture, GIF/PNG outputs and legend
	return ErrTODO
}
