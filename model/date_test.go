package model_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/model"
)

func TestDateJSON(t *testing.T) {
	d := model.Today()

	j, err := d.MarshalJSON()
	if err != nil {
		t.Fatal("marshal JSON", err.Error())
	}

	e := fmt.Sprintf(`"%s"`, d)
	if e != string(j) {
		t.Errorf("date JSON marshal\nEXP: %s\nACT: %s\n", e, j)
	}
}
