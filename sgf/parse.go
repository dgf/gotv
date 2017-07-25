package sgf

import (
	"bytes"
	"regexp"
	"strings"
	"text/scanner"
)

var (
	lineFeeds  = []string{"\r\n", "\n\r", "\r"}
	spaceRunes = regexp.MustCompile(`[ \t\f]+`)
	spaceList  = regexp.MustCompile(`[ ]+`)
	trimLines  = regexp.MustCompile(`([ ]*\n[ ]*)+`)
)

// clean SGF line feeds
func clean(sgf string) string {
	// replace all line feed combinations with a single one
	for _, f := range lineFeeds {
		sgf = strings.Replace(sgf, f, "\n", -1)
	}

	// combine multiple space to a single one
	return spaceRunes.ReplaceAllString(sgf, " ")
}

// property value scan (loop until ']')
func scanValue(s *scanner.Scanner) (v string) {
	b := bytes.Buffer{}

	for {
		r := s.Scan()
		if r == scanner.EOF || r == ']' {
			break // this is the end
		}
		switch r {
		case scanner.Ident:
			b.WriteString(s.TokenText())
		case scanner.Comment:
			b.WriteString(s.TokenText())
		case '\\':
			b.WriteRune(r)       // write \
			if s.Peek() == ']' { // check escaped end "\]"
				s.Scan() // write ]
				b.WriteString(s.TokenText())
			}
		default:
			b.WriteString(s.TokenText())
		}
	}

	v = b.String()
	v = spaceList.ReplaceAllString(v, " ")
	v = trimLines.ReplaceAllString(v, "\n")
	v = strings.Trim(v, " \n")
	return
}

// Parse SGF
func Parse(sgf string) Collection {
	var (
		collection Collection
		tree       *Tree
		node       Node
		ident      string
	)

	// setup scanner (only indent and comments match)
	s := &scanner.Scanner{}
	s.Init(strings.NewReader(clean(sgf)))
	s.Mode = scanner.ScanIdents | scanner.ScanComments

	var r rune
	for {
		if r == scanner.EOF || s.Peek() == scanner.EOF {
			break // this is the end
		}

		r = s.Scan()
		switch r {

		case '(': // sequence
			if tree == nil { // root
				tree = &Tree{}
				collection = append(collection, tree)
			} else { // subtree > go down
				t := &Tree{Parent: tree}
				tree.Collection = append(tree.Collection, t)
				tree = t
			}

		case ';': // node
			node = Node{}
			tree.Sequence = append(tree.Sequence, node)

		case ')': // end > go one up
			node = nil
			tree = tree.Parent
		}

		// property ident
		if node != nil && r == scanner.Ident {
			ident = s.TokenText()

			// skip until [
			for {
				if s.Peek() == scanner.EOF {
					break // this is the end
				}
				if s.Scan() == '[' {
					break
				}
			}

			if "C" == ident {
				s.Whitespace = 0 // skip nothing in comments
			} else {
				s.Whitespace = 1 << '\n' // skip line feeds in value
			}

			// scan property value
			node[ident] = scanValue(s)

			// reset white spaces ignore
			s.Whitespace = scanner.GoWhitespace
		}
	}

	return collection
}
