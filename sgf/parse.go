package sgf

import (
	"regexp"
	"strings"
)

var (
	collectionCleanUp = regexp.MustCompile(`(?m)[ \n]*([A-Z]+)[ \n]*\[[ \n]*((\\\]|[^\]]+)+)*\]`)
	spaceCleanUp      = regexp.MustCompile(`[ \t\f]+`)
)

// cleanup and parse SGF
//go:generate nex -s parse.nex
func Parse(s string) Collection {
	// replace all line feed combinations with a single one
	for _, r := range []string{"\r\n", "\n\r", "\r"} {
		s = strings.Replace(s, r, "\n", -1)
	}

	// combine multiple space to a single one
	s = spaceCleanUp.ReplaceAllString(s, " ")

	// remove all spaces and line feeds from property definitions
	s = collectionCleanUp.ReplaceAllString(s, "${1}[${2}]")

	// decode precleaned string
	return parse(s)
}
