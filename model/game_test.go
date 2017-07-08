package model_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/dgf/gotv/model"
)

func TestGameJSON(t *testing.T) {
	g := model.Game{Black: model.Player{Name: "icke", Rank: model.Kyu5}}
	e := fmt.Sprintf(`"black":{"name":"%s","rank":"%s"`, g.Black.Name, g.Black.Rank)
	if j, err := json.Marshal(g); err != nil {
		t.Fatal(err)
	} else if !strings.Contains(string(j), e) {
		t.Errorf("invalid JSON: %s\n", j)
	}
}
