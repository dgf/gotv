package model

// Game details
type Game struct {
	Date     Date     `json:"date"`
	Place    string   `json:"place"`
	Name     string   `json:"name"`
	Black    Player   `json:"black"`
	White    Player   `json:"white"`
	Size     Size     `json:"size"`
	Handicap Handicap `json:"handicap"`
	Komi     Komi     `json:"komi"`
	Result   string   `json:"result"`
}

// GameNames slice with names of games
func GameNames(games []Game) (names []string) {
	for _, g := range games {
		names = append(names, g.Name)
	}
	return
}
