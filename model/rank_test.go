package model_test

import (
	"testing"

	"github.com/dgf/gotv/model"
)

func TestRank(t *testing.T) {
	// assert unknown decode = NR
	u := model.ToRank("unknown")
	if u != model.NR {
		t.Errorf("unknown = NR expected, actual: %s\n", u)
	}

	// check forward and backward mappings
	for r, e := range map[model.Rank]string{
		model.NR:    "NR",
		model.Kyu30: "30 Kyu",
		model.Kyu17: "17 Kyu",
		model.Kyu4:  "4 Kyu",
		model.Kyu1:  "1 Kyu",
		model.Dan1:  "1 Dan",
		model.Dan9:  "9 Dan",
		model.Pro1:  "1 Pro",
		model.Pro9:  "9 Pro",
	} {
		// to string equals?
		a := r.String()
		if e != a {
			t.Errorf("%#v.String()\nEXP: %#v\nACT: %#v\n", r, e, a)
		}

		// from string equals?
		if r != model.ToRank(a) {
			t.Errorf("ToRank(%#v)\nEXP: %#v\nACT: %#v\n", a, r, model.ToRank(a))
		}
	}
}
