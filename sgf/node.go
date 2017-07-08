package sgf

import (
	"bytes"
	"fmt"

	"github.com/dgf/gotv/utils"
)

type Node struct {
	Properties map[string]string
}

// stringified node with properties sorted by ident (encode SGF Node)
func (n Node) String() string {
	s := bytes.Buffer{}
	s.WriteString(";")

	utils.SortAndCall(n.Properties, func(k, v string) {
		s.WriteString(fmt.Sprintf("%s[%s]", k, v))
	})

	return s.String()
}
