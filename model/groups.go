package model

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// a group on the board
type Group struct {
	Color     Color
	Liberties map[Point]struct{}
	Members   map[Point]struct{}
}

// stringified group dump with sorted members and liberties
// format: (B|W) m[members...] l[liberties...]
func (g Group) String() string {
	s := bytes.Buffer{}

	// color
	s.WriteString(fmt.Sprintf("%s", g.Color))

	// members
	members := []string{}
	for m := range g.Members {
		members = append(members, fmt.Sprintf("%v", m))
	}
	sort.Strings(members)
	s.WriteString(fmt.Sprintf(" m%v", members))

	// liberties
	liberties := []string{}
	for l := range g.Liberties {
		liberties = append(liberties, fmt.Sprintf("%v", l))
	}
	sort.Strings(liberties)
	s.WriteString(fmt.Sprintf(" l%v", liberties))

	return s.String()
}

// groups by ID
type Groups map[int]Group

// stringified groups dump sorted by ID
// format: ID (B|W) m[members...] l[liberties...]
func (g Groups) String() string {

	// sort group IDs
	ids := []int{}
	for id := range g {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	// stringify groups by id
	s := []string{}
	for _, id := range ids {
		s = append(s, fmt.Sprintf("%2d %s", id, g[id]))
	}

	return strings.Join(s, "\n")
}
