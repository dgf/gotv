package model_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/model/testdata"
)

var allFilter = model.Filter{}.
	DateMax(testdata.Tomorrow).DateMin(testdata.Yesterday).
	HandicapMax(model.H9).HandicapMin(model.H0).
	RankMax(model.Pro9).RankMin(model.NR).
	SizeMax(model.X19).SizeMin(model.X9)

func TestFilterMatches(t *testing.T) {
	for _, check := range []struct {
		name   string
		games  []string
		filter model.Filter
	}{
		{"date max", []string{"g1"}, model.Filter{}.DateMax(testdata.Yesterday)},
		{"date min max low", []string{"g1", "g2"}, model.Filter{}.DateMin(testdata.Yesterday).DateMax(testdata.Today)},
		{"date min max high", []string{"g2", "g3"}, model.Filter{}.DateMin(testdata.Today).DateMax(testdata.Tomorrow)},
		{"date min", []string{"g3"}, model.Filter{}.DateMin(testdata.Tomorrow)},

		{"handicap max", []string{"g1"}, model.Filter{}.HandicapMax(model.H0)},
		{"handicap min max low", []string{"g1", "g2"}, model.Filter{}.HandicapMin(model.H0).HandicapMax(model.H2)},
		{"handicap min max high", []string{"g2", "g3"}, model.Filter{}.HandicapMin(model.H2).HandicapMax(model.H3)},
		{"handicap min", []string{"g3"}, model.Filter{}.HandicapMin(model.H3)},

		{"rank max", []string{"g1"}, model.Filter{}.RankMax(model.Kyu5)},
		{"rank min max low", []string{"g1", "g2"}, model.Filter{}.RankMin(model.Kyu6).RankMax(model.Dan1)},
		{"rank min max high", []string{"g2", "g3"}, model.Filter{}.RankMin(model.Kyu2).RankMax(model.Pro1)},
		{"rank min", []string{"g3"}, model.Filter{}.RankMin(model.Dan7)},

		{"size max", []string{"g1"}, model.Filter{}.SizeMax(model.X9)},
		{"size min max low", []string{"g1", "g2"}, model.Filter{}.SizeMin(model.X9).SizeMax(model.X13)},
		{"size min max high", []string{"g2", "g3"}, model.Filter{}.SizeMin(model.X13).SizeMax(model.X19)},
		{"size min", []string{"g3"}, model.Filter{}.SizeMin(model.X19)},

		{"all", []string{"g1", "g2", "g3"}, allFilter},
	} {
		// filter
		list := []model.Game{}
		for _, g := range testdata.Games {
			if check.filter.Matches(&g) {
				list = append(list, g)
			}
		}

		// count
		if len(list) != len(check.games) {
			t.Errorf("invalid filter count: %s\n\tEXP: %d\n\tACT: %d\n", check.name, len(check.games), len(list))
		}

		// assert
		act := fmt.Sprintf("%s", model.GameNames(list))
		exp := fmt.Sprintf("%s", check.games)
		if exp != act {
			t.Errorf("invalid filter result: %s\n\tEXP: %s\n\tACT: %s\n", check.name, exp, act)
		}
	}
}

func BenchmarkFilterAll(b *testing.B) {
	g := &model.Game{
		Name: "test", Date: testdata.Today, Size: model.X13, Handicap: model.H2, Komi: 0.5,
		White: model.Player{Name: "icke", Rank: model.Kyu5},
		Black: model.Player{Name: "er", Rank: model.Kyu7},
	}
	for n := 0; n < b.N; n++ {
		allFilter.Matches(g)
	}
}
