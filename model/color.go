package model

// Color black or white (are no colors, but how to name it?)
type Color int

// color constants
const (
	Black Color = iota
	White
)

func (c Color) String() string {
	if c == Black {
		return "B"
	}
	return "W"
}
