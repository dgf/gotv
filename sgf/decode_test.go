package sgf_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/sgf"
)

func TestDecode(t *testing.T) {
	for game, properties := range map[model.Game]map[string]string{ // value
		{Name: "date", Date: model.ToDate("2017-07-08")}:        {"DT": "2017-07-08"},
		{Name: "place", Place: "Berlin"}:                        {"PL": "Berlin"},
		{Name: "handicap", Handicap: model.H2}:                  {"HA": "2"},
		{Name: "komi", Komi: 0.5}:                               {"KM": "0.5"},
		{Name: "size", Size: model.X9}:                          {"SZ": "9"},
		{Name: "result", Result: "B+R"}:                         {"RE": "B+R"},
		{Name: "b name", Black: model.Player{Name: "icke"}}:     {"PB": "icke"},
		{Name: "b rank", Black: model.Player{Rank: model.Kyu5}}: {"BR": "5 Kyu"},
		{Name: "b over", Black: model.Player{Name: "icke"}}:     {"PB": "icke", "BT": "team"},
		{Name: "b team", Black: model.Player{Name: "team"}}:     {"BT": "team"},
		{Name: "w name", White: model.Player{Name: "er"}}:       {"PW": "er"},
		{Name: "w rank", White: model.Player{Rank: model.Kyu3}}: {"WR": "3 Kyu"},
		{Name: "w over", White: model.Player{Name: "er"}}:       {"PW": "er", "WT": "team"},
		{Name: "w team", White: model.Player{Name: "team"}}:     {"WT": "team"},
	} {
		properties["GN"] = game.Name // sync name
		tree := sgf.Tree{Sequence: []*sgf.Node{&sgf.Node{Properties: properties}}}
		act := fmt.Sprintf("%s", tree.Decode()) // String() encodes it back
		exp := fmt.Sprintf("%s", game)
		if exp != act {
			t.Errorf("decode %s\nEXP: %s\nACT: %s\n", properties, exp, act)
		}
	}
}
