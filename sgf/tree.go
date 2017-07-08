package sgf

import (
	"bytes"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/utils"
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

func (gt *Tree) Decode() (g model.Game) {
	if len(gt.Sequence) < 1 {
		return
	}

	// map root node sequence (order matters for name overrides)
	utils.SortAndCall(gt.Sequence[0].Properties, func(k, v string) {
		if d, ok := decoderMap[k]; ok {
			d(&g, v)
		}
	})

	return
}
