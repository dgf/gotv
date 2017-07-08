package sgf_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/sgf"
)

func TestDecode(t *testing.T) {
	for g, p := range map[model.Game]map[string]string{
		{Date: model.ToDate("2017-07-08"), Place: "Berlin"}:   {"DT": "2017-07-08", "PL": "Berlin"},
		{Name: "test", Handicap: model.H2, Komi: 0.5}:         {"GN": "test", "HA": "2", "KM": "0.5"},
		{Size: model.X9, Result: "B+R"}:                       {"SZ": "9", "RE": "B+R"},
		{Black: model.Player{Name: "icke", Rank: model.Kyu5}}: {"BR": "5 Kyu", "PB": "icke"},
		{Black: model.Player{Name: "icke"}}:                   {"PB": "icke", "BT": "team"},
		{Black: model.Player{Name: "team"}}:                   {"BT": "team"},
		{White: model.Player{Name: "er", Rank: model.Kyu3}}:   {"WR": "3 Kyu", "PW": "er"},
		{White: model.Player{Name: "er"}}:                     {"PW": "er", "WT": "team"},
		{White: model.Player{Name: "team"}}:                   {"WT": "team"},
	} {
		a := sgf.Tree{Sequence: []*sgf.Node{&sgf.Node{Properties: p}}}
		gs := fmt.Sprintf("%s", g)
		ts := fmt.Sprintf("%s", a.Decode())
		if gs != ts {
			t.Errorf("decode: %s, \nEXP: %s\nACT: %s\n", p, gs, ts)
		}
	}
}
