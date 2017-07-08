package sgf

import (
	"bytes"
)

type Tree struct {
	Parent     *Tree
	Sequence   []*Node
	Collection []*Tree
}

// stringified game tree (encode SGF GameTree)
func (t Tree) String() string {
	s := bytes.Buffer{}
	s.WriteString("(")
	for _, n := range t.Sequence {
		s.WriteString(n.String())
	}
	for _, c := range t.Collection {
		s.WriteString(c.String())
	}
	s.WriteString(")")
	return s.String()
}
