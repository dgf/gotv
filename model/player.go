package model

import "fmt"

// Player name and rank
type Player struct {
	Name string `json:"name"`
	Rank Rank   `json:"rank"`
}

// String player name and rank
func (g Player) String() string {
	return fmt.Sprintf(`%s (%s)`, g.Name, g.Rank)
}

// MinRank of black or white
func MinRank(b Player, w Player) Rank {
	if b.Rank < w.Rank {
		return b.Rank
	}
	return w.Rank
}

// MaxRank of black or white
func MaxRank(b Player, w Player) Rank {
	if b.Rank > w.Rank {
		return b.Rank
	}
	return w.Rank
}
