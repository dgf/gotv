package sgf

import (
	"bytes"
)

// Collection game trees
type Collection []*Tree

// String tree collection (encode SGF Collection)
func (c Collection) String() string {
	s := bytes.Buffer{}
	for _, t := range c {
		s.WriteString(t.String())
	}
	return s.String()
}
