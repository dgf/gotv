package model

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	// reverse column lookup
	header = "+abcdefghjklmnopqrst"
)

var (
	// IGS compatible placement like "R16"
	rIGS = regexp.MustCompile("^[A-HJ-T]1?[0-9]$")
	// SGF compatible placement like "rd"
	rSGF = regexp.MustCompile("^[a-hj-t][a-hj-t]$")
)

// Board supports a straight forward Go game with capturing and KO + suicide detection
type Board struct {
	Size     int
	Move     int
	Last     Point
	KO       *Point
	Groups   Groups          // id > group
	Members  map[Point]int   // point > group
	Stones   map[Point]Color // point > color
	Captures struct {        // count
		Black int // white stones
		White int // black stones
	}
}

// NewBoard returns empty board
func NewBoard(s Size) Board {
	return Board{
		Size:    int(s),
		Groups:  Groups{},
		Members: map[Point]int{},
		Stones:  map[Point]Color{},
	}
}

// Place a stone, return captures
func (b *Board) Place(c Color, p Point) (Points, error) {

	// check border
	if p.X < 1 || p.X > b.Size || p.Y < 1 || p.Y > b.Size {
		return nil, fmt.Errorf("%v is out", p)
	}

	// check KO
	if b.KO != nil && b.KO.X == p.X && b.KO.Y == p.Y {
		return nil, fmt.Errorf("%v is KO", p)
	}

	// check intersection
	if _, ok := b.Stones[p]; ok {
		return nil, fmt.Errorf("%v is occupied", p)
	}

	// unique references
	liberties := map[Point]struct{}{}
	friends := map[int]Group{}
	opposites := map[int]Group{}
	captures := map[int]Group{}

	// categories neighbours
	for _, n := range neighbours(b.Size, p) {
		s, ok := b.Stones[n]
		if !ok { // empty
			liberties[n] = struct{}{}
			continue
		}
		id := b.Members[n]
		g := b.Groups[id]
		if s == c { // friendly
			friends[id] = g
		} else { // opposite
			if len(g.Liberties) == 1 { // taken by placement
				captures[id] = g
			} else {
				opposites[id] = g
			}
		}
	}

	// single stone group
	g := Group{
		Color:     c,
		Liberties: liberties,
		Stones:    map[Point]struct{}{p: {}},
	}

	// merge friends
	if len(friends) > 0 {
		for _, f := range friends {
			for l := range f.Liberties {
				g.Liberties[l] = struct{}{}
			}
			for m := range f.Stones {
				g.Stones[m] = struct{}{}
			}
		}
		// remove actual stone from merged liberties
		delete(g.Liberties, p)
	}

	// check suicide
	if len(captures) == 0 && len(g.Liberties) == 0 {
		return nil, fmt.Errorf("%v is suicide", p)
	}

	// add group
	b.Move++
	b.Last = p
	b.Groups[b.Move] = g
	b.Stones[p] = c

	// delete merged groups
	for id := range friends {
		delete(b.Groups, id)
	}

	// reference the new one
	for m := range g.Stones {
		b.Members[m] = b.Move
	}

	// reduce opposite liberties
	for _, o := range opposites {
		delete(o.Liberties, p)
	}

	// remove captures
	removed := Points{}
	for id, o := range captures {
		delete(b.Groups, id)
		for m := range o.Stones {
			delete(b.Members, m)
			delete(b.Stones, m)
			removed = append(removed, m)
		}
	}
	sort.Sort(removed)

	// recalc group liberties
	for _, r := range removed {
		for _, n := range neighbours(b.Size, r) {
			if id, ok := b.Members[n]; ok {
				b.Groups[id].Liberties[r] = struct{}{}
			}
		}
	}

	// sum captures
	if c == Black {
		b.Captures.Black += len(removed)
	} else {
		b.Captures.White += len(removed)
	}

	// detect KO
	b.KO = nil // one stone with only the captured liberty?
	if len(removed) == 1 && len(g.Stones) == 1 && len(g.Liberties) == 1 {
		b.KO = &removed[0] // mark capture as KO point
	}

	return removed, nil
}

// PlaceSGF stone
func (b *Board) PlaceSGF(c Color, p string) ([]string, error) {
	if !rSGF.MatchString(p) {
		return nil, fmt.Errorf("%s is invalid", p)
	}
	s := []string{}

	// decode and place
	r, err := b.Place(c, Point{
		X: strings.IndexByte(header, p[0]),
		Y: strings.IndexByte(header, p[1]),
	})

	// encode captures
	for _, p := range r {
		s = append(s, fmt.Sprintf("%s%s", string(header[p.X]), string(header[p.Y])))
	}

	return s, err
}

// PlaceIGS stone
func (b *Board) PlaceIGS(c Color, i string) ([]string, error) {
	if !rIGS.MatchString(i) {
		return nil, fmt.Errorf("%s is invalid", i)
	}
	s := []string{}

	// decode and place
	y, _ := strconv.Atoi(string(i[1:]))
	r, err := b.Place(c, Point{
		X: strings.Index(header, strings.ToLower(string(i[0]))),
		Y: b.Size + 1 - y, // reverse
	})

	// encode captures
	for _, p := range r {
		s = append(s, fmt.Sprintf("%s%d", strings.ToUpper(string(header[p.X])), b.Size+1-p.Y))
	}

	return s, err
}

func (b Board) String() string {
	s := bytes.Buffer{}
	s.WriteString(header[:b.Size+1])
	s.WriteString("+\n")

	for y := 1; y <= b.Size; y++ {
		s.WriteByte(header[y])
		for x := 1; x <= b.Size; x++ {
			if c, ok := b.Stones[Point{X: x, Y: y}]; !ok {
				if b.KO != nil && b.KO.X == x && b.KO.Y == y {
					s.WriteString("+")
				} else {
					s.WriteString(" ")
				}
			} else if c == Black {
				s.WriteString("●")
			} else {
				s.WriteString("◯")
			}
		}
		s.WriteString("|\n")
	}

	s.WriteString(fmt.Sprintf("%v w%d b%d", b.Last, b.Captures.White, b.Captures.Black))
	return s.String()
}
