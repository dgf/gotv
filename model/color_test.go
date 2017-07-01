package model_test

import (
	"testing"

	"github.com/dgf/gotv/model"
)

// short color aliases
const (
	B = model.Black
	W = model.White
)

func TestColorString(t *testing.T) {
	if "B" != B.String() || "W" != W.String() {
		t.Errorf("color string")
	}
}
