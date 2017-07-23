package model

import "sort"

// LessGame partial game comparison
type LessGame func(g1, g2 *Game) bool

// GameSorter sort games
type GameSorter []LessGame

// GameOrder returns a Sorter that sorts using the less functions, in order.
func GameOrder(less ...LessGame) GameSorter {
	return less
}

// ByDate newest first
func ByDate(g1, g2 *Game) bool {
	return g1.Date.After(g2.Date)
}

// ByDateDesc oldest first
func ByDateDesc(g1, g2 *Game) bool {
	return g1.Date.Before(g2.Date)
}

// ByHandicap lowest first
func ByHandicap(g1, g2 *Game) bool {
	return g1.Handicap < g2.Handicap
}

// ByHandicapDesc highest first
func ByHandicapDesc(g1, g2 *Game) bool {
	return g1.Handicap > g2.Handicap
}

// ByRank highest first
func ByRank(g1, g2 *Game) bool {
	return MaxRank(g1.Black, g1.White) > MaxRank(g2.Black, g2.White)
}

// ByRankDesc lowest first
func ByRankDesc(g1, g2 *Game) bool {
	return MinRank(g1.Black, g1.White) < MinRank(g2.Black, g2.White)
}

// BySize greatest first
func BySize(g1, g2 *Game) bool {
	return g1.Size > g2.Size
}

// BySizeDesc smallest first
func BySizeDesc(g1, g2 *Game) bool {
	return g1.Size < g2.Size
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (s *GameSorter) Sort(games []Game) {
	sort.Sort(&gameSorter{games: games, less: *s})
}

// gameSorter implements the Sort interface, sorting the games within.
type gameSorter struct {
	games []Game
	less  []LessGame
}

// Len is part of sort.Interface.
func (gs *gameSorter) Len() int {
	return len(gs.games)
}

// Swap is part of sort.Interface.
func (gs *gameSorter) Swap(i, j int) {
	gs.games[i], gs.games[j] = gs.games[j], gs.games[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or !Less.
func (gs *gameSorter) Less(i, j int) bool {
	g1, g2 := &gs.games[i], &gs.games[j]
	var l int // try all but the last
	for l = 0; l < len(gs.less)-1; l++ {
		less := gs.less[l]
		switch {
		case less(g1, g2):
			return true
		case less(g2, g1):
			return false
		}
	}
	// call last comparison
	return gs.less[l](g1, g2)
}
