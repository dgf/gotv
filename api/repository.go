package api

import (
	"github.com/dgf/gotv/model"
)

// Page results
type Page struct {
	Start int
	Count int
}

// Repository games
type Repository interface {

	// Add returns ID of stored game
	Add(g model.Game) (int, error)

	// Game returns game
	Game(int) (model.Game, error)

	// Games returns game count
	Games() (int, error)

	// List filtered and sorted games page
	List(f model.Filter, s model.GameSorter, p Page) ([]model.Game, error)
}
