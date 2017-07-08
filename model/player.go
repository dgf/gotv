package model

import "fmt"

type Player struct {
	Name string `json:"name"`
	Rank Rank   `json:"rank"`
}

// stringify
func (g Player) String() string {
	return fmt.Sprintf(`%s (%s)`, g.Name, g.Rank)
}
