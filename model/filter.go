package model

type flag int

const (
	minDate flag = 1 << iota
	maxDate
	minHand
	maxHand
	minRank
	maxRank
	minSize
	maxSize
)

// Filter games between min and max values
type Filter struct {
	flags   flag
	minDate Date
	maxDate Date
	minHand Handicap
	maxHand Handicap
	minRank Rank
	maxRank Rank
	minSize Size
	maxSize Size
}

// DateMin filters minimum Date
func (f Filter) DateMin(d Date) Filter {
	f.flags |= minDate
	f.minDate = d
	return f
}

// DateMax filters maximum Date
func (f Filter) DateMax(d Date) Filter {
	f.flags |= maxDate
	f.maxDate = d
	return f
}

// HandicapMin filters minimum Handicap
func (f Filter) HandicapMin(h Handicap) Filter {
	f.flags |= minHand
	f.minHand = h
	return f
}

// HandicapMax filters maximum Handicap
func (f Filter) HandicapMax(h Handicap) Filter {
	f.flags |= maxHand
	f.maxHand = h
	return f
}

// RankMin filters minimum Rank
func (f Filter) RankMin(r Rank) Filter {
	f.flags |= minRank
	f.minRank = r
	return f
}

// RankMax filters maximum Rank
func (f Filter) RankMax(r Rank) Filter {
	f.flags |= maxRank
	f.maxRank = r
	return f
}

// SizeMin filters minimum Size
func (f Filter) SizeMin(s Size) Filter {
	f.flags |= minSize
	f.minSize = s
	return f
}

// SizeMax filters maximum Size
func (f Filter) SizeMax(s Size) Filter {
	f.flags |= maxSize
	f.maxSize = s
	return f
}

func (f *Filter) active(a flag) bool {
	return f.flags&a != 0
}

// Matches all active filters
func (f *Filter) Matches(g *Game) bool {
	switch {
	case f.active(minDate) && g.Date.Before(f.minDate):
		fallthrough
	case f.active(maxDate) && g.Date.After(f.maxDate):
		fallthrough
	case f.active(minHand) && g.Handicap < f.minHand:
		fallthrough
	case f.active(maxHand) && g.Handicap > f.maxHand:
		fallthrough
	case f.active(minSize) && g.Size < f.minSize:
		fallthrough
	case f.active(maxSize) && g.Size > f.maxSize:
		fallthrough
	case f.active(minRank) && MaxRank(g.Black, g.White) < f.minRank:
		fallthrough
	case f.active(maxRank) && MinRank(g.Black, g.White) > f.maxRank:
		return false
	}
	return true
}
