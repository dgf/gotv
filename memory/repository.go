package memory

import (
	"errors"
	"sync"

	"github.com/dgf/gotv/api"
	"github.com/dgf/gotv/model"
)

type repository struct {
	mutex sync.RWMutex
	games []model.Game
}

// New creates a fresh in-memory game repository
func New() api.Repository {
	return &repository{
		games: []model.Game{},
	}
}

// Add is part of api.Repository
func (r *repository) Add(g model.Game) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.games = append(r.games, g)
	return len(r.games), nil
}

// Game is part of api.Repository
func (r *repository) Game(id int) (g model.Game, err error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if id < 0 || id > len(r.games)-1 {
		err = errors.New("out of range")
	} else {
		g = r.games[id-1]
	}
	return
}

// Games is part of api.Repository
func (r *repository) Games() (int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return len(r.games), nil
}

// List is part of api.Repository
func (r *repository) List(f model.Filter, s model.GameSorter, p api.Page) ([]model.Game, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	filtered := []model.Game{}
	for _, g := range r.games {
		if f.Matches(&g) {
			filtered = append(filtered, g)
		}
	}

	s.Sort(filtered)

	// empty
	if p.Start < 1 || p.Count < 1 || len(filtered) < p.Start {
		return []model.Game{}, nil
	}

	// the rest
	if len(filtered) < p.Start+p.Count {
		return filtered[p.Start-1:], nil
	}

	// page it
	return filtered[p.Start-1 : p.Start-1+p.Count], nil
}
