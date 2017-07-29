package memory_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/api"
	"github.com/dgf/gotv/memory"
	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/model/testdata"
)

func TestMemory(t *testing.T) {
	// init repository
	repo := memory.New()
	for _, game := range testdata.Games {
		repo.Add(game)
	}

	// count
	count, err := repo.Games()
	if err != nil {
		t.Error(err)
	}
	if count != 3 {
		t.Errorf("count: %d\n", count)
	}

	// get game
	game, err := repo.Game(1)
	if err != nil {
		t.Error(err)
	}
	if game.Name != "g1" {
		t.Errorf("g1 isn't first: %#v", game)
	}

	// list checks
	for _, check := range []struct {
		filter model.Filter
		sorter model.GameSorter
		page   api.Page
		list   []string
	}{
		{ // empty filter, sorter, page => empty result
			model.Filter{}, model.GameSorter{}, api.Page{}, []string{},
		},
		{ // empty filter, sort by date, cut page top => two games
			model.Filter{}, model.GameSorter{model.ByDate}, api.Page{Start: 1, Count: 2}, []string{"g3", "g2"},
		},
		{ // empty filter, sort by date, cut page bottom => two games
			model.Filter{}, model.GameSorter{model.ByDate}, api.Page{Start: 2, Count: 10}, []string{"g2", "g1"},
		},
		{ // filter date, sort reverse by date, cut page => two games
			model.Filter{}.DateMax(testdata.Today), model.GameSorter{model.ByDateDesc}, api.Page{Start: 1, Count: 3}, []string{"g1", "g2"},
		},
	} {
		games, err := repo.List(check.filter, check.sorter, check.page)
		if err != nil {
			t.Error(err)
		}
		act := fmt.Sprintf("%s", model.GameNames(games))
		exp := fmt.Sprintf("%s", check.list)
		if exp != act {
			t.Errorf("list diff\n\tEXP: %s\n\tACT: %s\n", exp, act)
		}
	}
}
