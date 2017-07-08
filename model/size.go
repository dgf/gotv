package model

import "fmt"

// game board size
type Size int

// size constants
const (
	X9  Size = 9
	X13 Size = 13
	X19 Size = 19
)

func (s Size) String() string {
	return fmt.Sprintf("%d", s)
}
