package model_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/dgf/gotv/model"
	"github.com/dgf/gotv/model/testdata"
)

func TestDateAfter(t *testing.T) {
	for _, check := range []struct {
		date  model.Date
		after model.Date
	}{
		{testdata.Yesterday, testdata.Tomorrow},
		{testdata.Yesterday, testdata.Today},
		{testdata.Today, testdata.Tomorrow},
	} {
		if !check.after.After(check.date) {
			t.Errorf("! %s < %s", check.date, check.after)
		}
	}
}

func TestDateBefore(t *testing.T) {
	for _, check := range []struct {
		date   model.Date
		before model.Date
	}{
		{testdata.Tomorrow, testdata.Yesterday},
		{testdata.Tomorrow, testdata.Today},
		{testdata.Today, testdata.Yesterday},
	} {
		if !check.before.Before(check.date) {
			t.Errorf("! %s > %s", check.date, check.before)
		}
	}
}

func TestDateJSON(t *testing.T) {
	d := model.Today()
	i := map[string]model.Date{"today": d}

	j, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}

	act := string(j)
	exp := fmt.Sprintf(`"%s"`, d)
	if !strings.Contains(act, exp) {
		t.Errorf("date JSON marshal\nEXP: contains %s\nACT: %s\n", exp, act)
	}
}
