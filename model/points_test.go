package model_test

import (
	"testing"

	"github.com/dgf/gotv/model"
)

func P(x, y int) model.Point {
	return model.Point{X: x, Y: y}
}

func TestDiffPoints(t *testing.T) {
	a := model.Points{P(1, 2), P(3, 4)}
	b := model.Points{P(4, 3), P(1, 2)}
	d := model.DiffPoints(a, b)
	if len(d) != 2 || d[0].X != 3 || d[0].Y != 4 || d[1].X != 4 || d[1].Y != 3 {
		t.Error("diff expected")
	}
}
