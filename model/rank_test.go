package model_test

import (
	"testing"

	"github.com/dgf/gotv/model"
)

func TestRankString(t *testing.T) {
	// assert unknown decode = NR
	unknown := model.ToRank("unknown")
	if unknown != model.NR {
		t.Errorf("unknown = NR expected, actual: %s\n", unknown)
	}

	// check forward and backward mappings
	for rank, exp := range map[model.Rank]string{
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
		act := rank.String()
		if exp != act {
			t.Errorf("%s.String()\nEXP: %#v\nACT: %#v\n", rank, exp, act)
		}

		// from string equals?
		if rank != model.ToRank(act) {
			t.Errorf("ToRank(%s)\nEXP: %#s\nACT: %#s\n", act, rank, model.ToRank(act))
		}
	}
}

func TestRankGreaterThan(t *testing.T) {
	for _, check := range []struct {
		rank    model.Rank
		greater model.Rank
	}{
		{model.Kyu30, model.NR},
		{model.Kyu1, model.Kyu30},
		{model.Kyu5, model.Kyu7},
		{model.Dan1, model.Kyu1},
		{model.Dan9, model.Dan1},
		{model.Pro1, model.Dan9},
		{model.Pro9, model.Pro1},
	} {
		if check.rank < check.greater {
			t.Errorf("! %s > %s\n", check.rank, check.greater)
		}
	}
}

func TestRankLessThan(t *testing.T) {
	for _, check := range []struct {
		rank model.Rank
		less model.Rank
	}{
		{model.Pro1, model.Pro9},
		{model.Dan9, model.Pro1},
		{model.Dan1, model.Dan9},
		{model.Kyu1, model.Dan1},
		{model.Kyu6, model.Kyu3},
		{model.Kyu30, model.Kyu1},
		{model.NR, model.Kyu30},
	} {
		if check.rank > check.less {
			t.Errorf("! %s < %s\n", check.rank, check.less)
		}
	}
}
