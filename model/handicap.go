package model

import "fmt"

// handicap of a game
type Handicap int

// handicap constants
const (
	H0 Handicap = iota
	_
	H2
	H3
	H4
	H5
	H6
	H7
	H8
	H9
)

func (h Handicap) String() string {
	return fmt.Sprintf("%d", h)
}
