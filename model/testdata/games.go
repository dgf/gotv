package testdata

import "github.com/dgf/gotv/model"

var Games = []model.Game{
	{
		Name: "g1", Date: Yesterday, Size: model.X9, Handicap: model.H0, Komi: 6.5,
		White: model.Player{Name: "icke", Rank: model.Kyu5},
		Black: model.Player{Name: "er", Rank: model.Kyu6},
	},
	{
		Name: "g2", Date: Today, Size: model.X13, Handicap: model.H2, Komi: 0.5,
		White: model.Player{Name: "dan", Rank: model.Dan1},
		Black: model.Player{Name: "kyu", Rank: model.Kyu2},
	},
	{
		Name: "g3", Date: Tomorrow, Size: model.X19, Handicap: model.H3, Komi: 0.5,
		White: model.Player{Name: "pro", Rank: model.Pro1},
		Black: model.Player{Name: "dan", Rank: model.Dan7},
	},
}
