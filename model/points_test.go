package model_test

import (
	"fmt"

	"github.com/dgf/gotv/model"
)

func P(x, y int) model.Point {
	return model.Point{X: x, Y: y}
}

func ExampleDiffPoints() {
	a := model.Points{P(1, 2), P(3, 4)}
	b := model.Points{P(4, 3), P(1, 2)}
	fmt.Println(model.DiffPoints(a, b))
	// Output: [{3 4} {4 3}]
}
