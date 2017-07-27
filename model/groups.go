package model

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// Group of stones on the board
type Group struct {
	Color     Color
	Liberties map[Point]struct{}
	Stones    map[Point]struct{}
}

// String group dump with sorted members and liberties
// format: (B|W) m[members...] l[liberties...]
func (g Group) String() string {
	s := bytes.Buffer{}

	// color
	s.WriteString(fmt.Sprintf("%s", g.Color))

	// stones
	stones := []string{}
	for m := range g.Stones {
		stones = append(stones, fmt.Sprintf("%v", m))
	}
	sort.Strings(stones)
	s.WriteString(fmt.Sprintf(" m%v", stones))

	// liberties
	liberties := []string{}
	for l := range g.Liberties {
		liberties = append(liberties, fmt.Sprintf("%v", l))
	}
	sort.Strings(liberties)
	s.WriteString(fmt.Sprintf(" l%v", liberties))

	return s.String()
}

// Groups map games by ID
type Groups map[int]Group

// String groups dump sorted by ID
// format: ID (B|W) m[members...] l[liberties...]
func (g Groups) String() string {

	// sort group IDs
	ids := []int{}
	for id := range g {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	// stringify groups by ID
	s := []string{}
	for _, id := range ids {
		s = append(s, fmt.Sprintf("%2d %s", id, g[id]))
	}

	return strings.Join(s, "\n")
}
