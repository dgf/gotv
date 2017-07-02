package sgf

import (
	"bytes"
	"fmt"
	"sort"
)

type Collection []*GameTree

// stringified game tree collection
func (c Collection) String() string {
	s := bytes.Buffer{}
	for _, g := range c {
		s.WriteString(fmt.Sprintf("(%s)", g))
	}
	return s.String()
}

type GameTree struct {
	Parent     *GameTree
	Sequence   []*Node
	Collection []*GameTree
}

// stringified game tree = streamlined SGF with sorted properties
func (gt GameTree) String() string {
	s := bytes.Buffer{}
	for _, n := range gt.Sequence {
		s.WriteString(fmt.Sprintf(";%s", n))
	}
	for _, g := range gt.Collection {
		s.WriteString(fmt.Sprintf("(%s)", g))
	}
	return s.String()
}

type Node struct {
	Properties map[string]string
}

// stringified node with properties sorted by ident
func (n Node) String() string {
	ids := []string{}
	for id, _ := range n.Properties {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	s := bytes.Buffer{}
	for _, id := range ids {
		s.WriteString(fmt.Sprintf("%s[%s]", id, n.Properties[id]))
	}
	return s.String()
}
