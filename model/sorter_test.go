package model_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/model/testdata"
)

var (
	unorderedGames = []model.Game{

		{Name: "ydx9h0nr", Date: testdata.Yesterday, Size: model.X9},
		{Name: "ydx3h0nr", Date: testdata.Yesterday, Size: model.X13},
		{Name: "ydx1h0nr", Date: testdata.Yesterday, Size: model.X19},

		{Name: "ydx1h9nr", Date: testdata.Yesterday, Size: model.X19, Handicap: model.H9},
		{Name: "ydx1h2nr", Date: testdata.Yesterday, Size: model.X19, Handicap: model.H2},

		{Name: "ydx1h0d1", Date: testdata.Yesterday, Size: model.X19, White: model.Player{Rank: model.Dan1}},
		{Name: "ydx1h0k1", Date: testdata.Yesterday, Size: model.X19, White: model.Player{Rank: model.Kyu1}},
		{Name: "ydx1h0k2", Date: testdata.Yesterday, Size: model.X19, White: model.Player{Rank: model.Kyu2}},

		{Name: "tyx9h0nr", Date: testdata.Today, Size: model.X9},
		{Name: "tyx3h0nr", Date: testdata.Today, Size: model.X13},
		{Name: "tyx1h0nr", Date: testdata.Today, Size: model.X19},

		{Name: "tyx1h9nr", Date: testdata.Today, Size: model.X19, Handicap: model.H9},
		{Name: "tyx1h2nr", Date: testdata.Today, Size: model.X19, Handicap: model.H2},

		{Name: "tyx1h0d1", Date: testdata.Today, Size: model.X19, White: model.Player{Rank: model.Dan1}},
		{Name: "tyx1h0k1", Date: testdata.Today, Size: model.X19, White: model.Player{Rank: model.Kyu1}},
		{Name: "tyx1h0k2", Date: testdata.Today, Size: model.X19, White: model.Player{Rank: model.Kyu2}},

		{Name: "tmx9h0nr", Date: testdata.Tomorrow, Size: model.X9},
		{Name: "tmx3h0nr", Date: testdata.Tomorrow, Size: model.X13},
		{Name: "tmx1h0nr", Date: testdata.Tomorrow, Size: model.X19},

		{Name: "tmx1h9nr", Date: testdata.Tomorrow, Size: model.X19, Handicap: model.H9},
		{Name: "tmx1h2nr", Date: testdata.Tomorrow, Size: model.X19, Handicap: model.H2},

		{Name: "tmx1h0d1", Date: testdata.Tomorrow, Size: model.X19, White: model.Player{Rank: model.Dan1}},
		{Name: "tmx1h0k1", Date: testdata.Tomorrow, Size: model.X19, White: model.Player{Rank: model.Kyu1}},
		{Name: "tmx1h0k2", Date: testdata.Tomorrow, Size: model.X19, White: model.Player{Rank: model.Kyu2}},
	}

	allAsc  = model.GameOrder(model.ByDate, model.BySize, model.ByRank, model.ByHandicap)
	allDesc = model.GameOrder(model.ByDateDesc, model.BySizeDesc, model.ByRankDesc, model.ByHandicapDesc)

	// sorted and reversed by init()
	sortedNames  = model.GameNames(unorderedGames)
	reverseNames = model.GameNames(unorderedGames)
)

func init() {
	sort.Strings(sortedNames)
	sort.Sort(sort.Reverse(sort.StringSlice(reverseNames)))
}

func TestSortOrder(t *testing.T) {
	for _, check := range []struct {
		name   string
		order  []string
		sorter model.GameSorter
		games  []model.Game
	}{
		{
			"by nothing", []string{"g1", "g2"}, model.GameSorter{}, []model.Game{{Name: "g1"}, {Name: "g2"}},
		},
		{
			"by date", []string{"tm", "ty", "yd"}, model.GameOrder(model.ByDate), []model.Game{
				{Name: "yd", Date: testdata.Yesterday},
				{Name: "tm", Date: testdata.Tomorrow},
				{Name: "ty", Date: testdata.Today}},
		}, {
			"by size", []string{"x19", "x13", "x9"}, model.GameOrder(model.BySize), []model.Game{
				{Name: "x13", Size: model.X13},
				{Name: "x9", Size: model.X9},
				{Name: "x19", Size: model.X19}},
		}, {
			"by rank", []string{"p9", "p1", "d9", "d1", "k1", "k9", "nr"}, model.GameOrder(model.ByRank), []model.Game{
				{Name: "nr", Black: model.Player{Rank: model.NR}},
				{Name: "k1", White: model.Player{Rank: model.Kyu1}},
				{Name: "k9", Black: model.Player{Rank: model.Kyu9}},
				{Name: "d1", Black: model.Player{Rank: model.Dan1}},
				{Name: "d9", White: model.Player{Rank: model.Dan9}},
				{Name: "p1", White: model.Player{Rank: model.Pro1}},
				{Name: "p9", Black: model.Player{Rank: model.Pro9}}},
		}, {
			"by handicap", []string{"h0", "h2", "h5", "h9"}, model.GameOrder(model.ByHandicap), []model.Game{
				{Name: "h2", Handicap: model.H2},
				{Name: "h9", Handicap: model.H9},
				{Name: "h5", Handicap: model.H5},
				{Name: "h0", Handicap: model.H0}},
		},
		{"ascending", sortedNames, allAsc, unorderedGames},
		{"descending", reverseNames, allDesc, unorderedGames},
	} {
		check.sorter.Sort(check.games)

		act := fmt.Sprintf("%s", model.GameNames(check.games))
		exp := fmt.Sprintf("%s", check.order)
		if exp != act {
			t.Errorf("invalid sorter result: %s\n\tEXP: %s\n\tACT: %s\n", check.name, exp, act)
		}
	}
}

func BenchmarkSortAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if n%2 == 0 {
			allAsc.Sort(unorderedGames)
		} else {
			allDesc.Sort(unorderedGames)
		}
	}
}
