package testdata

import "github.com/dgf/gotv/model"

var (
	Today = model.Today()

	Tomorrow = Today.Add(0, 0, 1)

	Yesterday = Today.Add(0, 0, -1)
)
