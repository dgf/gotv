package sgf

import (
	"bytes"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/utils"
)

// Tree with parent reference, node sequence and subtree slice
type Tree struct {
	Parent     *Tree
	Sequence   []*Node
	Collection []*Tree
}

// String game tree (encode SGF GameTree)
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

// Decode game tree
func (t *Tree) Decode() (g model.Game) {
	if len(t.Sequence) < 1 {
		return
	}

	// map root node sequence (order matters for name overrides)
	utils.SortAndCall(t.Sequence[0].Properties, func(k, v string) {
		if d, ok := decoder[k]; ok {
			d(&g, v)
		}
	})

	return
}
